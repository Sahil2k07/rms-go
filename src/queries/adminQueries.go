package queries

import "database/sql"

type AdminQueries struct {
	db *sql.DB
}

func NewAdminQueries(db *sql.DB) *AdminQueries {
	return &AdminQueries{db: db}
}

func (aq *AdminQueries) CreateNewJob(title, description, companyName string, userId int) (int64, error) {
	query := `
		INSERT INTO Job (title, description, companyName, postedBy)
		VALUES (?, ?, ?, ?)
	`

	result, err := aq.db.Exec(query, title, description, companyName, userId)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return id, nil
}
