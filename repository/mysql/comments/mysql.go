package comments

import (
	"EzMusix/bussiness/comments"
	"errors"

	"gorm.io/gorm"
)

type CommentsRepo struct {
	DBConn *gorm.DB
}

func NewCommentsRepo(db *gorm.DB) comments.Repository {
	return &CommentsRepo{
		DBConn: db,
	}
}

func (repo *CommentsRepo) Insert(commentsDomain comments.Domain) (comments.Domain, error) {
	rec := FromDomain(commentsDomain)
	if err := repo.DBConn.Create(&rec).Error; err != nil {
		return comments.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (repo *CommentsRepo) Get(commentsDomain comments.Domain) ([]comments.Domain, error) {
	rec := []Comments{}
	if err := repo.DBConn.Debug().Find(&rec).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []comments.Domain{}, errors.New("record not found")
		}
		return []comments.Domain{}, err
	}
	var domainPlaylist []comments.Domain
	for _, val := range rec {
		domainPlaylist = append(domainPlaylist, val.toDomain())
	}
	return domainPlaylist, nil
}

func (repo *CommentsRepo) Delete(commentsDomain comments.Domain) (comments.Domain, error) {
	rec := FromDomain(commentsDomain)
	if err := repo.DBConn.Where("id=?", commentsDomain.Id).First(&rec).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return comments.Domain{}, errors.New("record not found")
		}
	}
	repo.DBConn.Where("id = ?", commentsDomain.Id).Delete(&rec)
	return rec.toDomain(), nil
}
