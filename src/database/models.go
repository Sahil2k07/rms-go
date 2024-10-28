package database

import "time"

type User struct {
	ID              *int   `json:"id" db:"id"`
	Email           string `json:"email" db:"email"`
	Address         string `json:"address" db:"address"`
	UserType        string `json:"userType" db:"userType"`
	Password        string `json:"password" db:"password"`
	ProfileHeadline string `json:"profileHeadline" db:"profileHeadline"`
}

type Profile struct {
	ID                *int    `json:"id" db:"id"`
	ResumeFileAddress *string `json:"resumeFileAddress" db:"resumeFileAddress"`
	Skills            *string `json:"skills" db:"skills"`
	Education         *string `json:"education" db:"education"`
	Experience        *string `json:"experience" db:"experience"`
	Name              string  `json:"name" db:"name"`
	Email             string  `json:"email" db:"email"`
	Phone             *string `json:"phone" db:"phone"`
}

type Job struct {
	ID              *int       `json:"id" db:"id"`
	Title           string     `json:"title" db:"title"`
	Description     string     `json:"description" db:"description"`
	CompanyName     string     `json:"companyName" db:"companyName"`
	PostedOn        *time.Time `json:"postedOn" db:"postedOn"`
	PostedBy        int        `json:"postedBy" db:"postedBy"`
	TotalApplicants int        `json:"totalApplicants" db:"totalApplicants"`
}
