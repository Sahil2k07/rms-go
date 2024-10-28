package middlewares

import (
	"net/http"

	"github.com/Sahil2k07/rms-go/src/utils"
)

func IsApplicant(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(UserContext).(*UserAuthDetails)

		if !ok || user == nil || user.UserType != "Applicant" {
			utils.UnAuthorized(w, "Applicants only route.")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(UserContext).(*UserAuthDetails)

		if !ok || user == nil || user.UserType != "Admin" {
			utils.UnAuthorized(w, "Admins only route.")
			return
		}

		next.ServeHTTP(w, r)
	})
}
