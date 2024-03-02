package handler

import (
	"log"
	"strconv"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/ArmanAmandykuly/gomidterm/pkg/entities/task"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks", PostTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", GetTaskById).Methods("GET")

	http.ListenAndServe(":8081", router)

	return router
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := task.GetTasks()
	if err != nil {
		http.Error(w, "Error getting users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	serializedTasks, err := json.Marshal(tasks)
	if(err != nil) {
		http.Error(w, "Error marshaling json", http.StatusInternalServerError);
	}
	w.Write(serializedTasks)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	taskId, err1 := strconv.Atoi(vars["id"])
	if(err1 != nil) {
		http.Error(w, "Error converting to integer", http.StatusInternalServerError)
	}

	_task, err := task.GetTaskById(taskId)
	if(err != nil) {
		http.Error(w, "Error proceeding the request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	serializedTask, err := json.Marshal(_task)
	w.Write(serializedTask)
}

func PostTask(w http.ResponseWriter, r *http.Request) {
	var newTask task.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	
	createdTask, err := task.SaveTask(newTask)
	if err != nil {
		log.Println("Error saving task:", err)
		http.Error(w, "Error saving task", http.StatusInternalServerError)
		return
	}

	// Encode the created task as JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTask)
}