package centralclient

import "time"

type Organization struct {
	Id            int             `json:"id"`
	Title         string          `json:"title"`
	ParentId      int             `json:"parent_id"`
	Path          string          `json:"path"`
	Realm         string          `json:"realm"`
	CreatedAt     *time.Time      `json:"created_at"`
	UpdatedAt     *time.Time      `json:"updated_at"`
	Organizations []*Organization `json:"organizations"`
}

type Application struct {
	Id           int           `json:"id"`
	Name         string        `json:"name"`
	WriteAccess  bool          `json:"write_access"`
	CreatedAt    *time.Time    `json:"created_at"`
	UpdatedAt    *time.Time    `json:"updated_at"`
	Organization *Organization `json:"organization"`
}
