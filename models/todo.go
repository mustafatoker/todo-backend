package models

type Todo struct {
	Id        string `json:"id"`
	Name      string `json:"name" validate:"min=3,max=255"`
	CreatedAt string `json:"created_at"`
}
