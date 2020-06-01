package saving

import "errors"

type Repository interface {
	SaveUser(u User) (string, bool, error)
}

type UserService struct {
	repository Repository
}

func NewUserService(repository Repository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) SaveUser(u User) (string, bool, error) {
	userID, status, err := s.repository.SaveUser(u)
	if err != nil {
		return "", false, err
	}

	if status == false {
		return "", status, errors.New("user could not saved")
	}

	return userID, status, nil

}