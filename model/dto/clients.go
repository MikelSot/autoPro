package dto

import "time"

type CommentClient struct {
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Comment   string    `gorm:"type:varchar(350); default:''; not null" json:"comment"`
}

type CommentClients []*CommentClient