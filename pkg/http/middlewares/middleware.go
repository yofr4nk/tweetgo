package refmiddlewares

import (
	"context"
	"tweetgo/pkg/domain"
)

type userFinder interface {
	FindUserExists(email string) (bool, error)
}

type getUser func(email string) (domain.User, error)

type userSaver interface {
	SaveUser(u domain.User) (bool, error)
}

type getUserFromCtx func(ctx context.Context) (domain.User, error)
type setUserToCtx func(ctx context.Context, u domain.User) context.Context
