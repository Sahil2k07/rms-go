package controllers

import (
	"net/http"

	"github.com/Sahil2k07/rms-go/src/services"
)

type JobController struct {
	service *services.JobService
}

func NewJobController(service *services.JobService) *JobController {
	return &JobController{service: service}
}

func (jc *JobController) GetJobs(w http.ResponseWriter, r *http.Request) {

}

func (jc *JobController) ApplyJobs(w http.ResponseWriter, r *http.Request) {

}
