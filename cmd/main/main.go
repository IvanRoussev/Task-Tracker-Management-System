package main

import (
	"fmt"
	"github.com/IvanRoussev/taskManager/pkg/routes"
	"github.com/gorilla/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	cors :=
		handlers.CORS(
			handlers.AllowedMethods([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
			handlers.AllowedHeaders([]string{"Content-Type"}),
			handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "Access-Control-Allow-Methods", "Access-Control-Allow-Headers"}),
			handlers.ExposedHeaders([]string{"Access-Control-Allow-Origin", "Access-Control-Allow-Methods", "Access-Control-Allow-Headers"}),
		)

	r.Use(cors)

	r.Path("/swagger.yaml").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "swagger.yaml")
	}))

	r.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger.yaml")))

	routes.RegisterTaskManagerRoutes(r)
	http.Handle("/", r)
	err := http.ListenAndServe("localhost:9010", r)

	if err != nil {
		fmt.Printf("Could start Server: %v", err)
	}

}
