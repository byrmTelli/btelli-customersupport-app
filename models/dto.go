package models

import "time"

// Auth

type LoginDTO struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	AccessToken string `json:"token"`
}

// Entities

type ComplaintDTO struct {
	ID          uint            `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	UserID      uint            `json:"user_id"`
	CategoryID  uint            `json:"category_id"`
	Status      ComplaintStatus `json:"status"`
}

type CategoryDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CommentDTO struct {
	ID          uint      `json:"id"`
	ComplaintID uint      `json:"complaint_id"`
	UserID      uint      `json:"user_id"`
	CommentText string    `json:"comment_text"`
	CreatedAt   time.Time `json:"createdAt"`
}
