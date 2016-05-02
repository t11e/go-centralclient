package centralclient

import "time"

type Organization struct {
	Id        string     `json:"id"`
	Title     string     `json:"title"`
	ParentId  string     `json:"parent_id"`
	Realm     string     `json:"realm"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Application struct {
	Id           string        `json:"id"`
	Name         string        `json:"name"`
	WriteAccess  bool          `json:"write_access"`
	CreatedAt    *time.Time    `json:"created_at"`
	UpdatedAt    *time.Time    `json:"updated_at"`
	Organization *Organization `json:"organization"`
}
