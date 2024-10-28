package controllers

import (
	"net/http"

	"github.com/Sahil2k07/rms-go/src/services"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) Signup(w http.ResponseWriter, r *http.Request) {

}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {

}

func (uc *UserController) UploadResume(w http.ResponseWriter, r *http.Request) {

}
