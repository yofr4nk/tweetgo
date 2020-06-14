package refmiddlewares

import "tweetgo/pkg/domain"

type userFinder interface {
	FindUserExists(email string) (bool, error)
}

type userSaver interface {
	SaveUser(u domain.User) (string, bool, error)
}
