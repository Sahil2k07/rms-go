package services

import "github.com/Sahil2k07/rms-go/src/queries"

type AdminService struct {
	queries *queries.AdminQueries
}

func NewAdminService(queries *queries.AdminQueries) *AdminService {
	return &AdminService{queries: queries}
}
