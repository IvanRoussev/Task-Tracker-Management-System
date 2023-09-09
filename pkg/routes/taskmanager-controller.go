package routes

import (
	"github.com/IvanRoussev/taskManager/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterTaskManagerRoutes = func(router *mux.Router) {
	router.HandleFunc("/task/", controllers.GetTask).Methods("GET")
	router.HandleFunc("/task/", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/task/{taskId}", controllers.GetTaskById).Methods("GET")
}
