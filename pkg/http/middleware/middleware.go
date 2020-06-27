package middleware

import (
	"context"
	"mime/multipart"
	"tweetgo/pkg/domain"
)

type getUser func(email string) (domain.User, error)
type updateUser func(u domain.User, ID string) (bool, error)
type getUserFromCtx func(ctx context.Context) (domain.User, error)
type setUserToCtx func(ctx context.Context, u domain.User) context.Context
type uploadFile func(filePathName string, file multipart.File, fileHeader *multipart.FileHeader) (string, error)
type userSaver interface {
	SaveUser(u domain.User) (bool, error)
}
type userFinder interface {
	FindUserExists(email string) (bool, error)
}
