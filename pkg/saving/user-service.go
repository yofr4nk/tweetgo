package saving

import (
	"errors"
	"tweetgo/pkg/domain"
)

type Repository interface {
	SaveUser(u domain.User) (bool, error)
}

type UserService struct {
	repository Repository
}

func NewUserService(repository Repository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) SaveUser(u domain.User) (bool, error) {
	status, err := s.repository.SaveUser(u)
	if err != nil {
		return false, err
	}

	if status == false {
		return status, errors.New("user could not saved")
	}

	return status, nil

}
