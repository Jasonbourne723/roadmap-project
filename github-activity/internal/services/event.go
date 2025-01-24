package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"roadmap/github-activity/internal/models"
)

var Srv = new(srv)

type srv struct {
}

func (s *srv) Get(userName string) ([]models.Event, error) {

	client := new(http.Client)
	url := fmt.Sprintf("https://api.github.com/users/%s/events", userName)
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	var result []models.Event
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
