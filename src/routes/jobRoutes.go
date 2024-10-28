package routes

import (
	"net/http"

	"github.com/Sahil2k07/rms-go/src/controllers"
	"github.com/Sahil2k07/rms-go/src/database"
	"github.com/Sahil2k07/rms-go/src/middlewares"
	"github.com/Sahil2k07/rms-go/src/queries"
	"github.com/Sahil2k07/rms-go/src/services"
)

func JobRoutes(router *http.ServeMux) {
	jobQueries := queries.NewJobQueries(database.DB)
	jobService := services.NewJobService(jobQueries)
	jobController := controllers.NewJobController(jobService)

	router.Handle("/jobs", middlewares.Auth(http.HandlerFunc(jobController.GetJobs)))

	router.Handle("/jobs/apply", middlewares.Auth(middlewares.IsApplicant(http.HandlerFunc(jobController.ApplyJobs))))
}
