package services

import "github.com/Sahil2k07/rms-go/src/queries"

type UserService struct {
	queries *queries.UserQueries
}

func NewUserService(queries *queries.UserQueries) *UserService {
	return &UserService{queries: queries}
}
