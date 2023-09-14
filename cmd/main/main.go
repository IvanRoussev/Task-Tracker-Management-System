package main

import (
	"fmt"
	"github.com/IvanRoussev/taskManager/pkg/routes"
	"github.com/gorilla/handlers"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/").Handler(httpSwagger.Handler(
		httpSwagger.URL("localhost:9010/docs/swagger.yaml")))

	cors :=
		handlers.CORS(
			handlers.AllowedMethods([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
			handlers.AllowedHeaders([]string{"Content-Type"}),
			handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "Access-Control-Allow-Methods", "Access-Control-Allow-Headers"}),
			handlers.ExposedHeaders([]string{"Access-Control-Allow-Origin", "Access-Control-Allow-Methods", "Access-Control-Allow-Headers"}),
		)

	r.Use(cors)

	routes.RegisterTaskManagerRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Server running on http://localhost:9010 \n")

	err := http.ListenAndServe("localhost:9010", r)

	if err != nil {
		fmt.Printf("Could start Server: %v", err)
	}

}
