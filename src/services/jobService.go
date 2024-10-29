package services

import (
	"time"

	"github.com/Sahil2k07/rms-go/src/database"
	"github.com/Sahil2k07/rms-go/src/queries"
)

type JobService struct {
	queries *queries.JobQueries
}

func NewJobService(queries *queries.JobQueries) *JobService {
	return &JobService{queries: queries}
}

func (js *JobService) GetAllJobs() ([]database.Job, error) {
	var jobs []database.Job

	rows, err := js.queries.GetAllJobs()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var job database.Job
		var postedOnBytes []byte

		err := rows.Scan(&job.ID, &job.Title, &job.Description, &job.CompanyName, &postedOnBytes, &job.PostedBy, &job.TotalApplicants)
		if err != nil {
			return nil, err
		}

		postedOn, err := time.Parse("2006-01-02 15:04:05", string(postedOnBytes))
		if err != nil {
			return nil, err
		}
		job.PostedOn = &postedOn

		jobs = append(jobs, job)
	}

	return jobs, nil
}

func (js *JobService) ApplyJob(jobId int) error {
	total, err := js.queries.Applicants(jobId)
	if err != nil {
		return err
	}

	err = js.queries.ApplyJob(jobId, total+1)

	return err
}
