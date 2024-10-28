package routes

import (
	"net/http"

	"github.com/Sahil2k07/rms-go/src/controllers"
	"github.com/Sahil2k07/rms-go/src/database"
	"github.com/Sahil2k07/rms-go/src/middlewares"
	"github.com/Sahil2k07/rms-go/src/queries"
	"github.com/Sahil2k07/rms-go/src/services"
)

func AdminRoutes(router *http.ServeMux) {
	adminQueries := queries.NewAdminQueries(database.DB)
	adminService := services.NewAdminService(adminQueries)
	adminController := controllers.NewAdminController(adminService)

	router.Handle("/admin/job", middlewares.Auth(middlewares.IsAdmin(http.HandlerFunc(adminController.CreateJob))))

	router.Handle("/admin/job/{job_id}", middlewares.Auth(middlewares.IsAdmin(http.HandlerFunc(adminController.JobDetails))))

	router.Handle("/admin/applicants", middlewares.Auth(middlewares.IsAdmin(http.HandlerFunc(adminController.GetUsers))))

	router.Handle("/admin/applicant/{applicant_id}", middlewares.Auth(middlewares.IsAdmin(http.HandlerFunc(adminController.ApplicantData))))
}
