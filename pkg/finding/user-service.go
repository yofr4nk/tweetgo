package finding

import "tweetgo/pkg/domain"

type Repository interface {
	FindUserExists(email string) (int64, error)
	FindUser(email string) (domain.User, error)
}

type UserService struct {
	repository Repository
}

func NewUserService(repository Repository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) FindUserExists(email string) (bool, error) {
	count, err := s.repository.FindUserExists(email)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil

}

func (s *UserService) GetUser(email string) (domain.User, error) {
	u, err := s.repository.FindUser(email)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil

}
