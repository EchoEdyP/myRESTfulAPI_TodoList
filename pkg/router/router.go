package router

import (
	"RESTfulAPI_todos/internal/handlers"
	"RESTfulAPI_todos/pkg/database"
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(conn *sql.DB, config *database.Config) *mux.Router {
	router := mux.NewRouter()

	// CREATE
	router.HandleFunc("/api.example.com/todolist/managed-todolist", func(w http.ResponseWriter, r *http.Request) {
		handlers.InsertTodos(w, r, config)
	}).Methods(http.MethodPost)
	// GET ALL
	router.HandleFunc("/api.example.com/todolists/managedAll-todolists", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllTodos(w, r, config)
	}).Methods(http.MethodGet)
	// GET ID
	router.HandleFunc("/api.example.com/todolist/managed-todolist/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetIdTodo(w, r, config)
	}).Methods(http.MethodGet)
	// UPDATE TD
	router.HandleFunc("/api.example.com/todolists/managed-todolists/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateTitleDescriptionTodos(w, r, config)
	}).Methods(http.MethodPut)
	// UPDATE S
	router.HandleFunc("/api.example.com/todolist/managedStatus-todolist/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateStatusTodos(w, r, config)
	}).Methods(http.MethodPatch)
	// DELETE
	router.HandleFunc("/api.example.com/todolist/manage-todolist/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteTodos(w, r, config)
	}).Methods(http.MethodDelete)

	return router
}
