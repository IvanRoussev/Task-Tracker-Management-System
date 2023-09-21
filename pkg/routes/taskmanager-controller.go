package routes

import (
	"github.com/IvanRoussev/taskManager/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterTaskManagerRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/task", controllers.GetTask).Methods("GET")
	router.HandleFunc("/api/task", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/api/task/{taskId}", controllers.GetTaskById).Methods("GET")
	router.HandleFunc("/api/task/{taskId}", controllers.DeleteTaskById).Methods("DELETE")
	router.HandleFunc("/api/task/{taskId}", controllers.EditTaskCompletedById).Methods("PUT")
	router.HandleFunc("/api/task/{taskId}", controllers.EditTaskDescriptionByID).Methods("PUT")
	router.HandleFunc("/api/task/{taskTitle}", controllers.DeleteTaskByTitle).Methods("DELETE")
}
