package finding

type Repository interface {
	FindUserExists(email string) (int64, error)
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
