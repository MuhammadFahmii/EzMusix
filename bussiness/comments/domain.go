package comments

import "time"

type Domain struct {
	Id        int
	Content   string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Insert(comments Domain) (Domain, error)
	Get(comments Domain) ([]Domain, error)
	Delete(comments Domain) (Domain, error)
}

type Repository interface {
	Insert(comments Domain) (Domain, error)
	Get(comments Domain) ([]Domain, error)
	Delete(comments Domain) (Domain, error)
}
