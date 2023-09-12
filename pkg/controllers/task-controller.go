package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/IvanRoussev/taskManager/pkg/config"
	"github.com/IvanRoussev/taskManager/pkg/models"
	"github.com/IvanRoussev/taskManager/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
)

var newTask models.Task

var db = config.GetDB()

func GetTask(w http.ResponseWriter, r *http.Request) {
	newTasks, err := models.GetAllTasks(db)

	if err != nil {
		fmt.Printf("Could not get tasks: %v", err)
	}

	res, _ := json.Marshal(newTasks)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		fmt.Printf("Could not get all Task as Requested: %v", err)
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	CreateTask := &models.Task{}

	utils.ParseBody(r, CreateTask)
	b := CreateTask.CreateTask(db)

	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)

	if err != nil {
		fmt.Printf("Could not create requested task: %v", err)
	}
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	fmt.Println(args)
	taskId := args["taskId"]
	fmt.Println(taskId)
	task := models.GetTaskById(db, taskId)

	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)

	if err != nil {
		fmt.Printf("Could not Fin Task by ID: %v error: %v", taskId, err)
	}

}

func DeleteTaskById(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	taskId := args["taskId"]

	task := models.DeleteTaskById(db, taskId)

	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

type UpdateTaskRequest struct {
	UpdateCompleted bool `json:"Completed"`
}

type UpdateDescription struct {
	UpdateDescription string `json:"Description"`
}

func EditTaskCompletedById(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	taskId := args["taskId"]

	var updateRequest UpdateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&updateRequest)

	if err != nil {
		fmt.Printf("Could not get new Description from request: %v", err)
	}

	completed := updateRequest.UpdateCompleted

	var task *models.Task

	task = task.EditTaskCompletionById(db, taskId, completed)

	res, _ := json.Marshal(task)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)

	if err != nil {
		fmt.Printf("Could not Send back new Updated Task: %v", err)
	}
}

func EditTaskDescriptionByID(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	taskId := args["taskId"]

	var updatedDescription UpdateDescription
	err := json.NewDecoder(r.Body).Decode(&updatedDescription)

	if err != nil {
		fmt.Printf("Could not get new Description from request: %v", err)
	}

	newDescription := updatedDescription.UpdateDescription

	var task *models.Task

	task = task.EditDescriptionById(db, taskId, newDescription)

	res, _ := json.Marshal(task)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)

	if err != nil {
		fmt.Printf("Could not write to server: %v", err)
	}
}
