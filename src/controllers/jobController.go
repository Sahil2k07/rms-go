package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sahil2k07/rms-go/src/services"
	"github.com/Sahil2k07/rms-go/src/utils"
)

type JobController struct {
	service *services.JobService
}

func NewJobController(service *services.JobService) *JobController {
	return &JobController{service: service}
}

func (jc *JobController) GetJobs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WrongMethod(w)
		return
	}

	jobs, err := jc.service.GetAllJobs()
	if err != nil {

		utils.InternalServerError(w, err.Error())
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "Fetched All JobPosts Successfully",
		"data":    jobs,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func (jc *JobController) ApplyJobs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WrongMethod(w)
		return
	}

	jobIdStr := r.URL.Query().Get("job_id")
	if jobIdStr == "" {
		utils.InvalidInput(w, "job_id missing")
		return
	}

	jobId, err := strconv.Atoi(jobIdStr)
	if err != nil {
		utils.InvalidInput(w, "invalid job_id format")
		return
	}

	err = jc.service.ApplyJob(int(jobId))
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "Successfully Applied for the Job",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(response)
}
