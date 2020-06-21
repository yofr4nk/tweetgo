package saving

import (
	"errors"
	"tweetgo/pkg/domain"
)

type Repository interface {
	SaveUser(u domain.User) (bool, error)
	UpdateUser(usrData domain.UserDataContainer, ID string) error
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
		return status, errors.New("user could not be saved")
	}

	return status, nil

}

func (s *UserService) UpdateUser(u domain.User, ID string) (bool, error) {
	userContainer := make(domain.UserDataContainer)
	usrToUpdate := fillUserDataContainer(u, userContainer)

	err := s.repository.UpdateUser(usrToUpdate, ID)
	if err != nil {
		return false, errors.New("user could not be updated")
	}

	return true, nil

}

func fillUserDataContainer(u domain.User, usrContainer domain.UserDataContainer) domain.UserDataContainer {
	if len(u.Name) > 0 {
		usrContainer["name"] = u.Name
	}

	if len(u.LastName) > 0 {
		usrContainer["lastname"] = u.LastName
	}

	if len(u.UserBirthday) > 0 {
		usrContainer["userbirthday"] = u.UserBirthday
	}

	if len(u.Email) > 0 {
		usrContainer["email"] = u.Email
	}

	if len(u.Avatar) > 0 {
		usrContainer["avatar"] = u.Avatar
	}

	if len(u.Banner) > 0 {
		usrContainer["banner"] = u.Banner
	}

	if len(u.Biography) > 0 {
		usrContainer["biography"] = u.Biography
	}

	if len(u.Location) > 0 {
		usrContainer["location"] = u.Location
	}

	if len(u.WebSite) > 0 {
		usrContainer["website"] = u.WebSite
	}

	return usrContainer
}
