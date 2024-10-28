package routes

import (
	"encoding/json"
	"net/http"
)

func AppRoutes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)

			response := map[string]interface{}{
				"success": false,
				"message": "Route not found",
			}

			json.NewEncoder(w).Encode(response)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := map[string]interface{}{
			"success": true,
			"message": "Server is Active",
		}

		json.NewEncoder(w).Encode(response)
	})

	UserRoutes(router)

	AdminRoutes(router)

	JobRoutes(router)

	return router
}
