package posts

import "time"

type Post struct {
	Id        int        `json:"id,omitempty"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Category  string     `json:"category"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Status    string     `json:"status"`
}
