package services

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"roadmap/url-shortening/internal/models"
	"time"
)

const (
	urlFileName = "url.json"
)

var UrlSvc = new(UrlService)

type UrlService struct{}

func (u *UrlService) Add(long string) (models.Url, error) {
	list, err := load()
	if err != nil {
		return models.Url{}, err
	}
	item, exist := first(list, long)
	if exist {
		return item, nil
	}
	short := base64.URLEncoding.EncodeToString([]byte(long))[:6]
	var id int
	if len(list) == 0 {
		id = 1
	} else {
		id = list[len(list)-1].Id + 1
	}
	t := models.Url{
		Short:     short,
		Long:      long,
		CreatedAt: time.Now(),
		Id:        id,
	}
	list = append(list, t)
	save(list)
	return t, nil
}

func (u *UrlService) Del(short string) error {
	list, err := load()
	if err != nil {
		return err
	}
	for i, v := range list {
		if v.Short == short {
			if i == 0 {
				list = list[1:]
			} else if i == len(list)-1 {
				list = list[:len(list)-1]
			} else {
				list = append(list[:i], list[i+1:]...)
			}
			if err := save(list); err != nil {
				return err
			}
			break
		}
	}
	return nil
}

func (u *UrlService) Get(short string) (models.Url, error) {
	list, err := load()
	if err != nil {
		return models.Url{}, err
	}
	for _, v := range list {
		if v.Short == short {
			return v, nil
		}
	}
	return models.Url{}, errors.New("url is not exist")
}

func (u *UrlService) Statistics(url string) {

}

func load() ([]models.Url, error) {
	f, err := os.OpenFile(urlFileName, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	var list []models.Url
	decoder.Decode(&list)
	return list, nil
}

func save(list []models.Url) error {
	f, err := os.OpenFile(urlFileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	encoder := json.NewEncoder(f)
	encoder.Encode(list)
	return nil
}

func first(l []models.Url, url string) (models.Url, bool) {
	for _, v := range l {
		if v.Long == url {
			return v, true
		}
	}
	return models.Url{}, false
}
