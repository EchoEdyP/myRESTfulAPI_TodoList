package handlers

import (
	"RESTfulAPI_todos/helper"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/sirupsen/logrus"

	"RESTfulAPI_todos/pkg/database"
	"RESTfulAPI_todos/pkg/model"
)

// DeleteTodos is a function that handles the DELETE request for the API, delete a todo from the database.
func DeleteTodos(w http.ResponseWriter, r *http.Request, config *database.Config) {

	conn, err := database.ConnectDB(config)
	if err != nil {
		helper.InternalServerError(w, err)
		logrus.Error(err)
	}
	defer conn.Close()

	// Get parameter id from URL
	params := mux.Vars(r)
	id, ok := params["id"]
	// Check if id exists in data
	if !ok {
		helper.BadRequest(w, errors.New(" id not found"))
		logrus.Info(ok)
		return
	}

	// Check if Todo exist in the database
	var count int
	if err := conn.QueryRow("SELECT COUNT(*) FROM TodoList WHERE id=?", id).Scan(&count); err != nil {
		helper.InternalServerError(w, err)
		logrus.WithError(err).Error("Failed to check Todo existence in the database")
		return
	}
	// Jika tidak ada baris yang memiliki id yang sama dengan id yang dikirimkan dalam request,
	if count == 0 {
		helper.NotFound(w, errors.New(" id not found in the db"))
		return
	}

	// Delete Todo from the database
	if _, err := conn.Exec("DELETE FROM TodoList WHERE id=?", id); err != nil {
		helper.InternalServerError(w, err)
		logrus.WithError(err).Error("Failed to delete Todo from the database")
		return
	}

	// Prepare response
	apiResponse := model.Response{
		Status:  http.StatusOK,
		Message: "Todo with id " + id + " has been deleted",
		Data:    nil,
	}
	logrus.Info("Delete Todo successfully")

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(apiResponse.Status)
	json.NewEncoder(w).Encode(apiResponse)
}
