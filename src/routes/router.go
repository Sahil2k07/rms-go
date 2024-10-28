package routes

import "net/http"

func AppRoutes() *http.ServeMux {
	router := http.NewServeMux()

	return router
}
