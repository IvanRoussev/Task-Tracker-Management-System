package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/IvanRoussev/taskManager/pkg/config"
	"github.com/IvanRoussev/taskManager/pkg/models"
	"github.com/IvanRoussev/taskManager/pkg/utils"
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

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	CreateTask := models.Task{}

	utils.ParseBody(r, CreateTask)
	b := CreateTask.CreateTask(db)

	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
