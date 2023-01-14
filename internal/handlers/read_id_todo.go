package handlers

import (
	"RESTfulAPI_todos/error_handling"
	"RESTfulAPI_todos/helper"
	"RESTfulAPI_todos/pkg/database"
	"RESTfulAPI_todos/pkg/model"
	"errors"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

// handler GetIdTodos merupakan fungsi yang menangani request GETBYID/SELECTBYID pada API yang ditujukan untuk MEMBACA todo dari database.
func GetIdTodo(w http.ResponseWriter, r *http.Request, db database.DBConn) {
	var todos model.Todos
	var arrTodos []model.Todos

	conn, err := db.Connect()
	if err != nil {
		error_handling.InternalServerError(w, err)
		logrus.Error(err)
	}
	defer conn.Close()

	// Set log level
	logrus.SetLevel(logrus.DebugLevel)

	// Ambil parameter id dari URL
	params := mux.Vars(r)
	id := params["id"]

	// Check if Todo exist in the database
	var count int
	if err := conn.QueryRow("SELECT COUNT(*) FROM TodoList WHERE id=?", id).Scan(&count); err != nil {
		error_handling.InternalServerError(w, err)
		logrus.WithError(err).Error("Failed to check Todo existence in the database")
		return
	}

	// Jika tidak ada baris yang memiliki id yang sama dengan id yang dikirimkan dalam request,
	if count == 0 {
		error_handling.NotFound(w, errors.New(" id not found in db"))
		return
	}

	// MenJalankan query SELECT dengan WHERE id = id yang diambil dari URL
	rows, err := conn.Query("SELECT id, title, description, status FROM TodoList WHERE id = ?", id)
	if err != nil {
		error_handling.InternalServerError(w, err)
		logrus.Error(err)
		return
	}

	//// Loop sebanyak data yang dibaca dari database.
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

	apiResponse := model.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    arrTodos,
	}
	logrus.Info("Read Id Todo successfully")

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(apiResponse.Status)
	helper.WriteToResponseBody(w, apiResponse)
}
