package controllers

import (
	"net/http"

	"github.com/Sahil2k07/rms-go/src/services"
)

type AdminController struct {
	service *services.AdminService
}

func NewAdminController(service *services.AdminService) *AdminController {
	return &AdminController{service: service}
}

func (ac *AdminController) CreateJob(w http.ResponseWriter, r *http.Request) {

}

func (ac *AdminController) JobDetails(w http.ResponseWriter, r *http.Request) {

}

func (ac *AdminController) GetUsers(w http.ResponseWriter, r *http.Request) {

}

func (ac *AdminController) ApplicantData(w http.ResponseWriter, r *http.Request) {

}
