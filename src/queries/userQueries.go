package queries

import (
	"database/sql"
	"errors"

	"github.com/Sahil2k07/rms-go/src/database"
)

type UserQueries struct {
	db *sql.DB
}

func NewUserQueries(db *sql.DB) *UserQueries {
	return &UserQueries{db: db}
}

func (q *UserQueries) CheckExistingUser(email string) (string, error) {
	var existingEmail string

	query := `
		SELECT email from User
		WHERE email = ?
	`
	err := q.db.QueryRow(query, email).Scan(&existingEmail)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", err
		}

		return "", err
	}

	return existingEmail, nil
}

func (q *UserQueries) InsertUser(email, password, userType, address, profileHeadline, name string) error {
	userQuery := `
		INSERT INTO User (email, password, userType, address, profileHeadline)
		VALUES (?, ?, ?, ?, ?)
	`

	_, err := q.db.Exec(userQuery, email, password, userType, address, profileHeadline)
	if err != nil {
		return err
	}

	profileQuery := `
		INSERT INTO Profile (name, email)
		VALUES (?, ?)
	`

	_, err = q.db.Exec(profileQuery, name, email)
	if err != nil {
		return err
	}

	return nil
}

func (q *UserQueries) GetUserDetails(email string) (database.User, error) {
	var userDetails database.User

	query := `
		SELECT id, email, userType, password, address, profileHeadline
		FROM User
		WHERE email = ?
	`

	err := q.db.QueryRow(query, email).Scan(
		&userDetails.ID,
		&userDetails.Email,
		&userDetails.UserType,
		&userDetails.Password,
		&userDetails.Address,
		&userDetails.ProfileHeadline,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return database.User{}, errors.New("email does not exist")
		}

		return database.User{}, err
	}

	return userDetails, nil
}

func (q *UserQueries) UpdateProfile(email, skills, education, experience, phone string) error {

	query := `
		UPDATE Profile
		SET skills = ?, education = ?, experience = ?, phone = ?
		WHERE email = ?
	`
	_, err := q.db.Exec(query, skills, education, experience, phone, email)
	if err != nil {
		return err
	}

	return nil
}
