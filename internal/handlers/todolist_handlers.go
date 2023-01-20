package handlers

import (
	"RESTfulAPI_todos/pkg/database"
	"net/http"
)

type TodoListHandlers interface {
	Create(w http.ResponseWriter, r *http.Request, config *database.Config) error
	ReadAll(w http.ResponseWriter, r *http.Request, config *database.Config) error
	ReadById(w http.ResponseWriter, r *http.Request, config *database.Config) error
	UpdateTitleAndDescription(w http.ResponseWriter, r *http.Request, config *database.Config) error
	UpdateStatus(w http.ResponseWriter, r *http.Request, config *database.Config) error
	Delete(w http.ResponseWriter, r *http.Request, config *database.Config) error
}
