package queries

import (
	"database/sql"
	"errors"
	"time"

	"github.com/Sahil2k07/rms-go/src/database"
)

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

func (aq *AdminQueries) AllUsers() (*sql.Rows, error) {
	query := `
		SELECT id, email, password, userType, address, profileHeadline
		FROM User
	`

	rows, err := aq.db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (aq *AdminQueries) UserData(id int) (database.User, error) {
	var user database.User

	query := `
		SELECT email, password, userType, address, profileHeadline
		FROM User
		WHERE id = ?
	`

	err := aq.db.QueryRow(query, id).Scan(&user.Email, &user.Password, &user.UserType, &user.Address, &user.ProfileHeadline)
	if err != nil {
		return database.User{}, err
	}

	return user, nil
}

func (aq *AdminQueries) UserProfileData(email string) (database.Profile, error) {
	var profile database.Profile

	query := `
		SELECT name, skills, education, experience, phone
		FROM Profile
		WHERE email = ?
	`

	err := aq.db.QueryRow(query, email).Scan(&profile.Name, &profile.Skills, &profile.Education, &profile.Experience, &profile.Phone)
	if err != nil {
		return database.Profile{}, err
	}

	return profile, nil
}

func (aq *AdminQueries) JobDetails(id int) (database.Job, error) {
	var data database.Job

	var postedOnByte []byte

	query := `
		SELECT title, description, companyName, totalApplicants, postedBy, postedOn
		FROM Job
		WHERE id = ?
	`

	err := aq.db.QueryRow(query, id).Scan(
		&data.Title,
		&data.Description,
		&data.CompanyName,
		&data.TotalApplicants,
		&data.PostedBy,
		&postedOnByte,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return database.Job{}, errors.New("invalid job_id")
		}

		return database.Job{}, err
	}

	postedOn, err := time.Parse("2006-01-02 15:04:05", string(postedOnByte))
	if err != nil {
		return database.Job{}, err
	}

	data.PostedOn = &postedOn

	return data, nil
}
