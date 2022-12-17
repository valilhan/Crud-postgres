package models

type Language struct {
	Id        int64  `json:"id"`
	Year      int8   `json:"year"`
	Name      string `json:"name"`
	Developer string `json:"developer"`
}
