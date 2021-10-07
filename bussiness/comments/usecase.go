package comments

import "errors"

type CommentsUsecase struct {
	commentsRepo Repository
}

func NewCommentsUsecase(repo Repository) Usecase {
	return &CommentsUsecase{
		commentsRepo: repo,
	}
}

func (CommentsUsecase *CommentsUsecase) Insert(comments Domain) (Domain, error) {
	res, err := CommentsUsecase.commentsRepo.Insert(comments)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}
func (CommentsUsecase *CommentsUsecase) Delete(comments Domain) (Domain, error) {
	res, err := CommentsUsecase.commentsRepo.Delete(comments)
	if err != nil {
		if err.Error() == "record not found" {
			return Domain{}, errors.New("record not found")
		}
		return Domain{}, err
	}
	return res, nil
}

func (CommentsUsecase *CommentsUsecase) Get(comments Domain) ([]Domain, error) {
	res, err := CommentsUsecase.commentsRepo.Get(comments)
	if err != nil {
		if err.Error() == "record not found" {
			return []Domain{}, errors.New("not found")
		}
		return []Domain{}, err
	}
	return res, nil
}
