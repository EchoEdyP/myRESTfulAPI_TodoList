package router

import (
	"RESTfulAPI_todos/internal/handlers"
	"RESTfulAPI_todos/pkg/database"
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(conn *sql.DB, config *database.Config, listHandlers handlers.TodoListHandlers) *mux.Router {
	router := mux.NewRouter()

	// CREATE
	router.HandleFunc("/api.example.com/todolist/managed-todolist", func(w http.ResponseWriter, r *http.Request) {
		listHandlers.Create(w, r, config)
	}).Methods(http.MethodPost)
	// GET ALL
	router.HandleFunc("/api.example.com/todolists/managedAll-todolists", func(w http.ResponseWriter, r *http.Request) {
		listHandlers.ReadAll(w, r, config)
	}).Methods(http.MethodGet)
	// GET ID
	router.HandleFunc("/api.example.com/todolist/managed-todolist/{id}", func(w http.ResponseWriter, r *http.Request) {
		listHandlers.ReadById(w, r, config)
	}).Methods(http.MethodGet)
	// UPDATE TD
	router.HandleFunc("/api.example.com/todolists/managed-todolists/{id}", func(w http.ResponseWriter, r *http.Request) {
		listHandlers.UpdateTitleAndDescription(w, r, config)
	}).Methods(http.MethodPut)
	// UPDATE S
	router.HandleFunc("/api.example.com/todolist/managedStatus-todolist/{id}", func(w http.ResponseWriter, r *http.Request) {
		listHandlers.UpdateStatus(w, r, config)
	}).Methods(http.MethodPatch)
	// DELETE
	router.HandleFunc("/api.example.com/todolist/manage-todolist/{id}", func(w http.ResponseWriter, r *http.Request) {
		listHandlers.Delete(w, r, config)
	}).Methods(http.MethodDelete)

	return router
}
