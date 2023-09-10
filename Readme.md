# Task Manager app

## *HTTP Routes*

### GET

* GetAllTasks()
  * This Route will get all Tasks in the database and return as http response


* GetTaskByID(id *string*)
  * This Route return Task with the Id provided in the http request 
    * Example  http://localhost:9010/{taskId}


### POST

* CreateTask() 
  * This is used by calling a Task struct before function call `ex. task1.createTask()`
  * This Route will create Task based on provided json in the body of request
  * example of json request body
    * `{
    "Description": "testing",
    "DueDate": "Aug 1",
    "Completed": false
    }`

### DELETE 

* DeleteTaskById(id *string*)
  * This function will delete a Task by Id provided in http request similar to EditTaskById


### PUT

* EditTaskCompletedById(id *string*)
  * This route Will edit the completion of the task 
  * send a request of the id in the header and a json object completion update to true or false
  * ex. req body `{ "completed" : true }`