package services

import "github.com/Sahil2k07/rms-go/src/queries"

type JobService struct {
	queries *queries.JobQueries
}

func NewJobService(queries *queries.JobQueries) *JobService {
	return &JobService{queries: queries}
}
