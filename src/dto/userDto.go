package dto

import "github.com/Sahil2k07/rms-go/src/database"

type SignupDto struct {
	Email           string            `json:"email" validate:"required,email,excludesall=(;)=:"`
	Password        string            `json:"password" validate:"required,excludesall=(;)=:"`
	UserType        database.UserType `json:"userType" validate:"required,excludesall=(;)=:"`
	Name            string            `json:"name" validate:"required,excludesall=(;)=:"`
	ProfileHeadline string            `json:"profileHeadline" validate:"required,excludesall=(;)=:"`
	Address         string            `json:"address" validate:"required,excludesall=;=:"`
}

type LoginDto struct {
	Email    string `json:"email" validate:"required,email,excludesall=(;)=:"`
	Password string `json:"password" validate:"required,excludesall=(;)=:"`
}

type ResumeAPIResponse struct {
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	Phone      string   `json:"phone"`
	Education  []any    `json:"education"`
	Experience []any    `json:"experience"`
	Skills     []string `json:"skills"`
}
