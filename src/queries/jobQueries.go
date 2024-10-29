package queries

import (
	"database/sql"
	"errors"
)

type JobQueries struct {
	db *sql.DB
}

func NewJobQueries(db *sql.DB) *JobQueries {
	return &JobQueries{db: db}
}

func (jq *JobQueries) GetAllJobs() (*sql.Rows, error) {
	query := `
		SELECT id, title, description, companyName, postedOn,postedBy , totalApplicants
		FROM Job
	`

	rows, err := jq.db.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (jq *JobQueries) Applicants(jobId int) (int, error) {
	var totalApplicants int

	query := `
		SELECT totalApplicants FROM Job
		WHERE id = ?
	`

	err := jq.db.QueryRow(query, jobId).Scan(&totalApplicants)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errors.New("no job with the mentioned job_id")
		}
		return 0, err
	}

	return totalApplicants, nil
}

func (jq *JobQueries) ApplyJob(jobId, total int) error {
	query := `
		UPDATE Job
		SET totalApplicants = ?
		WHERE id = ?
	`

	_, err := jq.db.Exec(query, total, jobId)

	return err
}
