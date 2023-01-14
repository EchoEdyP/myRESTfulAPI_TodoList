package router

import (
	"RESTfulAPI_todos/internal/handlers"
	"RESTfulAPI_todos/pkg/database"
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter returns a new instance of a mux router with the appropriate routes configured
func NewRouter(conn *database.MySQLConn) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api.example.com/todolist/managed-todolist", func(w http.ResponseWriter, r *http.Request) {
		handlers.InsertTodos(w, r, conn)
	})
	router.HandleFunc("/api.example.com/todolists/managedAll-todolists", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllTodos(w, r, conn)
	})

	router.HandleFunc("/api.example.com/todolist/managed-todolist/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetIdTodo(w, r, conn)
	})

	router.HandleFunc("/api.example.com/todolists/managed-todolists", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateTitleDescriptionTodos(w, r, conn)
	})

	router.HandleFunc("/api.example.com/todolist/managedStatus-todolist", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateStatusTodos(w, r, conn)
	})

	router.HandleFunc("/api.example.com/todolist/manage-todolist", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteTodos(w, r, conn)
	})

	return router
}
