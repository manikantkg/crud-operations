package routes

import (
	"net/http"
	"project/controller"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Router() {
	newRoute := mux.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{""},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPut,
		},
		AllowedHeaders:   []string{""},
		AllowCredentials: false,
	})

	handler := cors.Handler(newRoute)

	newRoute.HandleFunc("/test", controller.TestingApi).Methods("GET")
	newRoute.HandleFunc("/user", controller.User).Methods("GET")

	http.ListenAndServe(":9090", handler)
}
