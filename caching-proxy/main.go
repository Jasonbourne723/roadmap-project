package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"roadmap/caching-proxy/srv"
)

func main() {

	port := flag.Int("port", 8000, "端口号")
	origin := flag.String("origin", "https://jasonbourne723.github.io", "源地址")
	clear := flag.Bool("clear", false, "清理缓存")
	flag.Parse()
	var proxy = srv.Proxy
	if *clear {
		proxy.Clear()
	}

	if len(*origin) == 0 {

		log.Fatal("origin is null")
	}
	proxy.Origin = *origin
	http.Handle("/", proxy)
	log.Println("Listen to 8000")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))

}
