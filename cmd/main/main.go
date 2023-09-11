package main

import (
	"fmt"
	"github.com/IvanRoussev/taskManager/pkg/routes"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterTaskManagerRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Server running on http://localhost:9010")

	err := http.ListenAndServe("localhost:9010", r)

	if err != nil {
		fmt.Printf("Could start Server: %v", err)
	}

}
