package request

import (
	"EzMusix/bussiness/comments"
	"time"
)

type Comments struct {
	Id        int       `json:"id"`
	Content   string    `json:"content" form:"content"`
	UserID    int       `json:"user_id" form:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteComments struct {
	Id int `param:"id"`
}

func ToDomain(c Comments) comments.Domain {
	return comments.Domain{
		Id:        c.Id,
		Content:   c.Content,
		UserID:    c.UserID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
func DeleteToDomain(pl DeleteComments) comments.Domain {
	return comments.Domain{
		Id: pl.Id,
	}
}
