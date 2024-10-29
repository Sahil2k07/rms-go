package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	fmt.Println(user.Id)

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

}

func (ac *AdminController) GetUsers(w http.ResponseWriter, r *http.Request) {

}

func (ac *AdminController) ApplicantData(w http.ResponseWriter, r *http.Request) {

}
