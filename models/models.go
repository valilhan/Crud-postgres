package models

type Language struct {
	Id        int64  `json:"id"`
	Year      int    `json:"year"`
	Name      string `json:"name"`
	Developer string `json:"developer"`
}
