package main

import (
	"fmt"
	"net/http"

	"github.com/IvanRoussev/taskManager/pkg/routes"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("C:/Users/ivanr/Documents/Personal/Projects/Go/taskManager/api/swagger/swagger.yaml")))

	routes.RegisterTaskManagerRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Server running on http://localhost:9010")

	err := http.ListenAndServe("localhost:9010", r)

	if err != nil {
		fmt.Printf("Could start Server: %v", err)
	}

}
