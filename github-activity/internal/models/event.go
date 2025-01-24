package models

import "time"

type Event struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Actor     Actor     `json:"actor"`
	Repo      Repo      `json:"repo"`
	Payload   Payload   `json:"payload"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
}
type Actor struct {
	ID           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
}
type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Payload struct {
	Ref          interface{} `json:"ref"`
	RefType      string      `json:"ref_type"`
	MasterBranch string      `json:"master_branch"`
	Description  interface{} `json:"description"`
	PusherType   string      `json:"pusher_type"`
}
