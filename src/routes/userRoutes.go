package routes

import (
	"net/http"

	"github.com/Sahil2k07/rms-go/src/controllers"
	"github.com/Sahil2k07/rms-go/src/database"
	"github.com/Sahil2k07/rms-go/src/middlewares"
	"github.com/Sahil2k07/rms-go/src/queries"
	"github.com/Sahil2k07/rms-go/src/services"
)

func UserRoutes(router *http.ServeMux) {
	userQueries := queries.NewUserQueries(database.DB)
	userService := services.NewUserService(userQueries)
	userController := controllers.NewUserController(userService)

	router.HandleFunc("/signup", userController.Signup)

	router.HandleFunc("/login", userController.Login)

	router.Handle("/uploadResume", middlewares.Auth(middlewares.IsApplicant(http.HandlerFunc(userController.UploadResume))))
}
