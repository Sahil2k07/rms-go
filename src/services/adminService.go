package services

import (
	"database/sql"
	"errors"

	"github.com/Sahil2k07/rms-go/src/database"
	"github.com/Sahil2k07/rms-go/src/queries"
)

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

func (as *AdminService) GetAllUsers() ([]database.User, error) {
	var users []database.User

	rows, err := as.queries.AllUsers()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user database.User

		err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.UserType, &user.Address, &user.ProfileHeadline)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (as *AdminService) ApplicantData(id int) (database.User, database.Profile, error) {
	data, err := as.queries.UserData(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return database.User{}, database.Profile{}, errors.New("invalid appliacnt_id")
		}
		return database.User{}, database.Profile{}, err
	}

	if data.UserType == "Admin" {
		return database.User{}, database.Profile{}, errors.New("applicant is an admin")
	}

	profileData, err := as.queries.UserProfileData(data.Email)
	if err != nil {
		return database.User{}, database.Profile{}, err
	}

	return data, profileData, nil
}

func (as *AdminService) JobData(id int) (database.Job, error) {

	data, err := as.queries.JobDetails(id)
	if err != nil {
		return database.Job{}, err
	}

	return data, nil
}
