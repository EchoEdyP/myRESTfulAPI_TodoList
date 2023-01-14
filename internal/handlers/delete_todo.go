package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"

	"RESTfulAPI_todos/error_handling"
	"RESTfulAPI_todos/pkg/database"
	"RESTfulAPI_todos/pkg/model"
)

// DeleteTodos is a function that handles the DELETE request for the API, delete a todo from the database.
func DeleteTodos(w http.ResponseWriter, r *http.Request, db database.DBConn) {

	conn, err := db.Connect()
	if err != nil {
		error_handling.InternalServerError(w, err)
		logrus.Error(err)
	}
	defer conn.Close()

	// Decode request body into map[string]interface{}
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		error_handling.BadRequest(w, err)
		logrus.WithError(err).Error("Failed to decode request body")
		return
	}

	// Check if id exists in data
	id, ok := data["id"]
	if !ok {
		error_handling.BadRequest(w, errors.New(" id not found"))
		logrus.Info(ok)
		return
	}

	// Check if Todo exist in the database
	var count int
	if err := conn.QueryRow("SELECT COUNT(*) FROM TodoList WHERE id=?", id).Scan(&count); err != nil {
		error_handling.InternalServerError(w, err)
		logrus.WithError(err).Error("Failed to check Todo existence in the database")
		return
	}

	if count == 0 {
		error_handling.NotFound(w, errors.New(" id not found in the db"))
		return
	}

	// Delete Todo from the database
	if _, err := conn.Exec("DELETE FROM TodoList WHERE id=?", id); err != nil {
		error_handling.InternalServerError(w, err)
		logrus.WithError(err).Error("Failed to delete Todo from the database")
		return
	}

	// Prepare response
	apiResponse := model.Response{
		Status:  http.StatusOK,
		Message: "Todo with id " + id.(string) + " has been deleted",
		Data:    nil,
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(apiResponse.Status)
	json.NewEncoder(w).Encode(apiResponse)
}
