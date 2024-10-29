package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sahil2k07/rms-go/src/dto"
	"github.com/Sahil2k07/rms-go/src/middlewares"
	"github.com/Sahil2k07/rms-go/src/services"
	"github.com/Sahil2k07/rms-go/src/utils"
	"github.com/go-playground/validator/v10"
)

type AdminController struct {
	service  *services.AdminService
	validate *validator.Validate
}

func NewAdminController(service *services.AdminService) *AdminController {
	return &AdminController{service: service, validate: validator.New()}
}

func (ac *AdminController) CreateJob(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WrongMethod(w)
		return
	}

	var createJobData dto.CreateJobDto

	err := json.NewDecoder(r.Body).Decode(&createJobData)
	if err != nil {
		utils.InvalidInput(w, "Invalid Json Format")
		return
	}

	if err := ac.validate.Struct(createJobData); err != nil {
		utils.InvalidInput(w, "Invalid Payload")
		return
	}

	user, ok := r.Context().Value(middlewares.UserContext).(*middlewares.UserAuthDetails)
	if !ok {
		utils.UnAuthorized(w, "User Credentials not found")
		return
	}

	id, err := ac.service.CreateJobPost(
		user.Id,
		createJobData.Title,
		createJobData.Description,
		createJobData.CompanyName,
	)
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "Job Post created successfully",
		"data": map[string]interface{}{
			"id":          id,
			"title":       createJobData.Title,
			"description": createJobData.Description,
			"companyName": createJobData.CompanyName,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}

func (ac *AdminController) JobDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WrongMethod(w)
		return
	}

	jobIdStr := r.PathValue("job_id")
	if jobIdStr == "" {
		utils.InvalidInput(w, "job_id missing")
		return
	}

	jobId, err := strconv.Atoi(jobIdStr)
	if err != nil {
		utils.InvalidInput(w, "invalid applicant_id format")
		return
	}

	data, err := ac.service.JobData(jobId)
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "Fetched job details",
		"data":    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func (ac *AdminController) GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WrongMethod(w)
		return
	}

	users, err := ac.service.GetAllUsers()
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}

	response := map[string]interface{}{
		"success": false,
		"message": "Fetched All Users Data Successfully",
		"data":    users,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func (ac *AdminController) ApplicantData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WrongMethod(w)
		return
	}

	userIdStr := r.PathValue("applicant_id")
	if userIdStr == "" {
		utils.InvalidInput(w, "applicant_id missing")
		return
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		utils.InvalidInput(w, "invalid applicant_id format")
		return
	}

	userData, profileData, err := ac.service.ApplicantData(userId)
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "Fetched Applicant Details",
		"data": map[string]interface{}{
			"user":    userData,
			"profile": profileData,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
