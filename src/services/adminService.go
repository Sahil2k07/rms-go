package services

import "github.com/Sahil2k07/rms-go/src/queries"

type AdminService struct {
	queries *queries.AdminQueries
}

func NewAdminService(queries *queries.AdminQueries) *AdminService {
	return &AdminService{queries: queries}
}

func (as *AdminService) CreateJobPost(id int, title, description, companyName string) (int64, error) {
	jobId, err := as.queries.CreateNewJob(title, description, companyName, id)
	if err != nil {
		return 0, err
	}

	return jobId, nil
}
