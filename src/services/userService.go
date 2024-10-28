package services

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Sahil2k07/rms-go/src/database"
	"github.com/Sahil2k07/rms-go/src/dto"
	"github.com/Sahil2k07/rms-go/src/queries"
	"github.com/Sahil2k07/rms-go/src/utils"
)

type UserService struct {
	queries *queries.UserQueries
}

func NewUserService(queries *queries.UserQueries) *UserService {
	return &UserService{queries: queries}
}

func (us *UserService) SignupUser(signupData dto.SignupDto) error {
	if signupData.UserType != "Admin" && signupData.UserType != "Applicant" {
		return errors.New("invalid userType field")
	}

	email, err := us.queries.CheckExistingUser(signupData.Email)
	if err != nil {
		return err
	}

	if email != "" {
		return errors.New("email already exists")
	}

	hashedPassword, err := utils.HashPassword(signupData.Password)
	if err != nil {
		return err
	}

	err = us.queries.InsertUser(
		signupData.Email,
		hashedPassword,
		string(signupData.UserType),
		signupData.Address,
		signupData.ProfileHeadline,
		signupData.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) LoginUser(loginData dto.LoginDto) (string, database.User, error) {
	data, err := us.queries.GetUserDetails(loginData.Email)
	if err != nil {
		return "", database.User{}, err
	}

	unhashedPassword := utils.CheckPasswordHash(loginData.Password, data.Password)
	if !unhashedPassword {
		return "", database.User{}, errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(*data.ID, data.Email, data.UserType)
	if err != nil {
		return "", database.User{}, err
	}

	return token, data, nil
}

func (us *UserService) UploadResume(file io.Reader, email string) (dto.ResumeAPIResponse, error) {
	var resumeResponse dto.ResumeAPIResponse

	apiUrl := os.Getenv("API_URL")
	apiKey := os.Getenv("API_KEY")

	if apiKey == "" || apiUrl == "" {
		return dto.ResumeAPIResponse{}, errors.New("api_key and api_secret not present in env")
	}

	// Create a POST request with binary body
	req, err := http.NewRequest("POST", apiUrl, file)
	if err != nil {
		return dto.ResumeAPIResponse{}, err
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("apikey", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return dto.ResumeAPIResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return dto.ResumeAPIResponse{}, errors.New("failed to parse resume: external API error")
	}

	if err := json.NewDecoder(res.Body).Decode(&resumeResponse); err != nil {
		return dto.ResumeAPIResponse{}, errors.New("failed to decode resume response")
	}

	// Extract relevant fields with fallback to "N/A"
	skills := "N/A"
	if len(resumeResponse.Skills) > 0 {
		skills = strings.Join(resumeResponse.Skills, ", ")
	}

	education := "N/A"
	if len(resumeResponse.Education) > 0 {
		eduJson, _ := json.Marshal(resumeResponse.Education) // Convert to JSON string
		education = string(eduJson)
	}

	experience := "N/A"
	if len(resumeResponse.Experience) > 0 {
		expJson, _ := json.Marshal(resumeResponse.Experience) // Convert to JSON string
		experience = string(expJson)
	}

	phone := "N/A"
	if resumeResponse.Phone != "" {
		phone = resumeResponse.Phone
	}

	err = us.queries.UpdateProfile(
		email,
		skills,
		string(education),
		string(experience),
		phone,
	)
	if err != nil {
		return dto.ResumeAPIResponse{}, err
	}

	return resumeResponse, nil
}
