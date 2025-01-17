package srv

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var Proxy = &proxy{
	cache: make(map[string]Cache, 5),
	lock:  &sync.RWMutex{},
}

type proxy struct {
	Origin string
	cache  map[string]Cache
	lock   *sync.RWMutex
}

type Cache struct {
	resp *http.Response
	body []byte
}

func (p *proxy) Clear() {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.cache = make(map[string]Cache)
}

func (p *proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	key := fmt.Sprintf("%s:%s", r.Method, r.URL.String())
	p.lock.RLock()
	if cache, ok := p.cache[key]; ok {
		p.lock.RUnlock()
		responseWithHeader(w, cache.resp, cache.body, true)
		fmt.Printf("hit cache: %v\n", key)
		return
	}
	p.lock.RUnlock()

	client := http.Client{}
	orginURL := p.Origin + r.URL.String()
	res, err := client.Get(orginURL)
	if err != nil {

	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	p.lock.Lock()
	p.cache[key] = Cache{
		resp: res,
		body: body,
	}
	p.lock.Unlock()
	responseWithHeader(w, res, body, false)
}

func responseWithHeader(writer http.ResponseWriter, res *http.Response, body []byte, isHit bool) {

	var hit string
	if isHit {
		hit = "HIT"
	} else {
		hit = "MISS"
	}
	writer.Header().Set("X-Cache", hit)
	for k, v := range res.Header {
		writer.Header()[k] = v
	}
	writer.Write(body)
}
