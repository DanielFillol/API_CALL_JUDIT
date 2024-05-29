package models

type BodyRequest struct {
	Search Search `json:"search"`
}

type Search struct {
	SearchType string `json:"search_type"`
	SearchKey  string `json:"search_key"`
}
