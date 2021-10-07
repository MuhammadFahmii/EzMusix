package response

import (
	"EzMusix/bussiness/comments"
	"time"
)

type Comments struct {
	Id        int       `json:"id" form:"id"`
	Content   string    `json:"content" form:"content"`
	UserID    int       `json:"user_id" form:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AddComments struct {
	Content string `json:"content"`
}
type DeleteComments struct {
	Content string
}

func FromDomain(pl comments.Domain) Comments {
	return Comments{
		Id:        pl.Id,
		Content:   pl.Content,
		UserID:    pl.UserID,
		CreatedAt: pl.CreatedAt,
		UpdatedAt: pl.UpdatedAt,
	}
}

func FromDomainDelete(pl comments.Domain) DeleteComments {
	return DeleteComments{
		Content: pl.Content,
	}
}

func ToAddComments(pl comments.Domain) AddComments {
	return AddComments{
		Content: pl.Content,
	}
}
