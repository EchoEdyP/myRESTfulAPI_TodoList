// Package handlers merupakan package yang berisi handler-handler yang akan digunakan pada API
package handlers

import (
	"RESTfulAPI_todos/helper"
	"RESTfulAPI_todos/pkg/database"
	"RESTfulAPI_todos/pkg/model"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// InsertTodos is a function that handles the CREATE/INSERT request for the API, adding a new todo to the database.
func InsertTodos(w http.ResponseWriter, r *http.Request, config *database.Config) {

	conn, err := database.ConnectDB(config)
	if err != nil {
		helper.InternalServerError(w, err)
		logrus.Error(err)
	}
	defer conn.Close()

	var todo model.Todos
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	validate := validator.New()
	err = validate.Struct(todo)
	if err != nil {
		helper.BadRequest(w, err)
		logrus.Error(err)
		return
	}

	res, err := conn.Exec("INSERT INTO TodoList(title, description) VALUES(?, ?)", todo.Title, todo.Description)
	if err != nil {
		helper.InternalServerError(w, err)
		logrus.Error(err)
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		logrus.Error(err)
		return
	}

	// Prepare response
	apiResponse := model.Response{
		Status:  http.StatusCreated,
		Message: "you have successfully created todo list with ID: " + strconv.FormatInt(lastID, 10),
		Data:    nil,
	}
	logrus.Info("Create Todo successfully")

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(apiResponse.Status)
	helper.WriteToResponseBody(w, apiResponse)

}
