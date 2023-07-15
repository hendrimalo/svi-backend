package posts

import "time"

type CreatePostInput struct {
	Title     string    `json:"title" binding:"required,min=20"`
	Content   string    `json:"content" binding:"required,min=200"`
	Category  string    `json:"category" binding:"required,min=3"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Status    string    `json:"status" binding:"oneof='publish' 'draft'"`
}

type UpdatePostInput struct {
	Title     string    `json:"title" binding:"required,min=20"`
	Content   string    `json:"content" binding:"required,min=200"`
	Category  string    `json:"category" binding:"required,min=3"`
	UpdatedAt time.Time `json:"created_at,omitempty"`
	Status    string    `json:"status" binding:"oneof='publish' 'draft'"`
}
