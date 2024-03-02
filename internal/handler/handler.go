package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"tsis2/pkg/entities/task"
	"tsis2/pkg/database/postgres"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Define your routes here
	router.HandleFunc("/tasks", GetAllTasks).Methods("GET")
	// router.HandleFunc("/posts", GetAllPosts).Methods("GET")

	return router
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	// Get all users from the database using the db package
	tasks, err := DBGetAllTasks()
	if err != nil {
		http.Error(w, "Error getting users", http.StatusInternalServerError)
		return
	}

	// Return the users as JSON
	// (You should implement a proper JSON response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("List of users:"))
	for _, task := range task {
		w.Write([]byte("\n" + user.Username))
	}
}

// GetAllPosts handles the /posts endpoint
// func GetAllPosts(w http.ResponseWriter, r *http.Request) {
// 	// Get all posts from the database using the db package
// 	posts, err := DBGetAllPosts()
// 	if err != nil {
// 		http.Error(w, "Error getting posts", http.StatusInternalServerError)
// 		return
// 	}

// 	// Return the posts as JSON
// 	// (You should implement a proper JSON response)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("List of posts:"))
// 	for _, post := range posts {
// 		w.Write([]byte("\n" + post.Title))
// 	}
// }