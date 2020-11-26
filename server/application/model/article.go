package model

type Article struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
