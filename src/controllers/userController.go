package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Sahil2k07/rms-go/src/dto"
	"github.com/Sahil2k07/rms-go/src/middlewares"
	"github.com/Sahil2k07/rms-go/src/services"
	"github.com/Sahil2k07/rms-go/src/utils"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	service  *services.UserService
	validate *validator.Validate
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		service:  service,
		validate: validator.New(),
	}
}

func (uc *UserController) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WrongMethod(w)
		return
	}

	var signupBody dto.SignupDto

	err := json.NewDecoder(r.Body).Decode(&signupBody)
	if err != nil {
		utils.InvalidInput(w, "Invalid Json Format")
		return
	}

	if err := uc.validate.Struct(signupBody); err != nil {
		utils.InvalidInput(w, "Invalid Payload")
		return
	}

	err = uc.service.SignupUser(signupBody)
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "Signup successfull",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WrongMethod(w)
		return
	}

	var loginBody dto.LoginDto

	err := json.NewDecoder(r.Body).Decode(&loginBody)
	if err != nil {
		utils.InvalidInput(w, "Invalid Json Format")
		return
	}

	if err := uc.validate.Struct(loginBody); err != nil {
		utils.InvalidInput(w, "Invalid Payload")
		return
	}

	token, data, err := uc.service.LoginUser(loginBody)
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   3 * 24 * 60 * 60, // 3 days
	})

	response := map[string]interface{}{
		"success": true,
		"message": "Login Successfull",
		"token":   token,
		"data": map[string]interface{}{
			"id":       data.ID,
			"email":    data.Email,
			"userType": data.UserType,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func (uc *UserController) UploadResume(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WrongMethod(w)
		return
	}

	file, header, err := r.FormFile("resume")
	if err != nil {
		utils.InternalServerError(w, "Failed to Parse File")
		return
	}
	defer file.Close()

	// Check file type (only allow PDF or DOCX)
	if !(header.Filename[len(header.Filename)-4:] == ".pdf" || header.Filename[len(header.Filename)-5:] == ".docx") {
		utils.InvalidInput(w, "File must be in PDF or DOCX format")
		return
	}

	user, ok := r.Context().Value(middlewares.UserContext).(*middlewares.UserAuthDetails)
	if !ok {
		utils.UnAuthorized(w, "User Credentials not found")
		return
	}

	data, err := uc.service.UploadResume(file, user.Email)
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "Resume Uploaded Successfully",
		"data": map[string]interface{}{
			"skills":     data.Skills,
			"education":  data.Education,
			"experience": data.Education,
			"phone":      data.Phone,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
