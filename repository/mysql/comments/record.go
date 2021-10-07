package comments

import (
	"EzMusix/bussiness/comments"
	"time"
)

type Comments struct {
	Id        int `gorm:"primaryKey"`
	Content   string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromDomain(comments comments.Domain) Comments {
	return Comments{
		Id:        comments.Id,
		Content:   comments.Content,
		UserID:    comments.UserID,
		CreatedAt: comments.CreatedAt,
		UpdatedAt: comments.UpdatedAt,
	}
}

func (c *Comments) toDomain() comments.Domain {
	return comments.Domain{
		Id:        c.Id,
		Content:   c.Content,
		UserID:    c.UserID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
