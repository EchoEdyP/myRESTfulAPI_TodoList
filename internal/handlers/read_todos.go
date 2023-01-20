package handlers

import (
	"RESTfulAPI_todos/helper"
	"RESTfulAPI_todos/pkg/database"
	"RESTfulAPI_todos/pkg/model"
	"github.com/sirupsen/logrus"
	"net/http"
)

// handler GetAllTodos merupakan fungsi yang menangani request GETALL/SELECTALL pada API yang ditujukan untuk MEMBACA todo dari database.
func GetAllTodos(w http.ResponseWriter, r *http.Request, config *database.Config) {
	var todos model.Todos
	var arrTodos []model.Todos

	conn, err := database.ConnectDB(config)
	if err != nil {
		helper.InternalServerError(w, err)
		logrus.Error(err)
	}
	defer conn.Close()

	// Set log level
	logrus.SetLevel(logrus.DebugLevel) // logrus.DebugLevel: Menampilkan semua log, termasuk log debug.

	// Get All Todo from the database
	rows, err := conn.Query("SELECT id, title, description, status FROM TodoList")
	if err != nil {
		helper.InternalServerError(w, err)
		logrus.Error(err)
		return
	}

	// Loop sebanyak data yang dibaca dari database.
	for rows.Next() {
		// Scan data yang dibaca dari database ke dalam variabel todos.
		err = rows.Scan(&todos.Id, &todos.Title, &todos.Description, &todos.Status)
		if err != nil {
			// Jika terjadi error, log error tersebut.
			logrus.Fatal(err)
		} else {
			// Jika tidak terjadi error, tambahkan data todo ke dalam slice arrTodos.
			arrTodos = append(arrTodos, todos)
		}
	}

	// Prepare response
	apiResponse := model.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    arrTodos,
	}
	logrus.Info("Read All Todo successfully")

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(apiResponse.Status)
	helper.WriteToResponseBody(w, apiResponse)
}
