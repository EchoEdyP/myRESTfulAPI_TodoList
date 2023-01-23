package handlers

import (
	"RESTfulAPI_todos/helper"
	"RESTfulAPI_todos/pkg/database"
	"RESTfulAPI_todos/pkg/model"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type TodoListhandlersImpl struct {
	DB *sql.DB
}

func NewTodoListhandlersImpl(DB *sql.DB) *TodoListhandlersImpl {
	return &TodoListhandlersImpl{DB: DB}
}

func (handlers *TodoListhandlersImpl) Create(w http.ResponseWriter, r *http.Request, config *database.Config) error {

	var todo model.Todos
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	validate := validator.New()
	err = validate.Struct(todo)
	if err != nil {
		helper.BadRequest(w, err)
		logrus.Error(err)
		return nil
	}

	res, err := handlers.DB.Exec("INSERT INTO TodoList(title, description) VALUES(?, ?)", todo.Title, todo.Description)
	if err != nil {
		helper.InternalServerError(w, err)
		logrus.Error(err)
		return nil
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		logrus.Error(err)
		return nil
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

	return nil
}

func (handlers *TodoListhandlersImpl) ReadAll(w http.ResponseWriter, r *http.Request, config *database.Config) error {
	var todos model.Todos
	var arrTodos []model.Todos

	// Set log level
	logrus.SetLevel(logrus.DebugLevel) // logrus.DebugLevel: Menampilkan semua log, termasuk log debug.

	// Get All Todo from the database
	rows, err := handlers.DB.Query("SELECT id, title, description, status FROM TodoList")
	if err != nil {
		helper.InternalServerError(w, err)
		logrus.Error(err)
		return nil
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

	return nil
}

func (handlers *TodoListhandlersImpl) ReadById(w http.ResponseWriter, r *http.Request, config *database.Config) error {
	var todos model.Todos
	var arrTodos []model.Todos

	// Set log level
	logrus.SetLevel(logrus.DebugLevel)

	// Get parameter id from URL
	params := mux.Vars(r)
	id := params["id"]

	// Check if Todo exist in the database
	// Menjalankan query untuk menghitung/mencari jumlah baris pada tabel TodoList yang memiliki id yang sama dengan id yang dikirimkan dalam request.
	var count int
	if err := handlers.DB.QueryRow("SELECT COUNT(*) FROM TodoList WHERE id=?", id).Scan(&count); err != nil {
		helper.InternalServerError(w, err)
		logrus.WithError(err).Error("Failed to check Todo existence in the database")
		return nil
	}

	// Jika tidak ada baris yang memiliki id yang sama dengan id yang dikirimkan dalam request,
	if count == 0 {
		helper.NotFound(w, errors.New(" id not found in db"))
		return nil
	}

	// Get Id Todo from the database
	rows, err := handlers.DB.Query("SELECT id, title, description, status FROM TodoList WHERE id = ?", id)
	if err != nil {
		helper.InternalServerError(w, err)
		logrus.Error(err)
		return nil
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

	// Prepare response
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

	return nil
}

func (handlers *TodoListhandlersImpl) UpdateTitleAndDescription(w http.ResponseWriter, r *http.Request, config *database.Config) error {

	// Membaca request yang dikirimkan dalam format raw JSON
	var request model.Todos
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		helper.InternalServerError(w, err)
		// Jika terjadi error saat mengambil data, log error tersebut
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error parsing raw JSON request")
		return nil
	}

	// Menyimpan data yang dibaca dari request ke dalam variabel id, title, dan description.
	params := mux.Vars(r)
	id := params["id"]
	title := request.Title
	description := request.Description

	// Check if Todo exist in the database
	// Menjalankan query untuk menghitung/mencari jumlah baris pada tabel TodoList yang memiliki id yang sama dengan id yang dikirimkan dalam request.
	var count int
	err = handlers.DB.QueryRow("SELECT COUNT(*) FROM TodoList WHERE id=?", id).Scan(&count)
	// Jika terjadi error saat menjalankan query, log error menggunakan logrus dan keluar dari fungsi.
	if err != nil {
		helper.InternalServerError(w, err)
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error checking id")
		return nil
	}

	// Jika tidak ada baris yang memiliki id yang sama dengan id yang dikirimkan dalam request,
	if count == 0 {
		helper.NotFound(w, err)
		return nil
	}

	/*validasi input menggunakan playground validator*/
	validate := validator.New()
	err = validate.Struct(model.Todos{
		Title:       title,
		Description: description,
	})

	// Jika terjadi error saat menjalankan validasi, log error menggunakan logrus
	if err != nil {
		helper.BadRequest(w, err)
		logrus.WithFields(logrus.Fields{"error": err}).Error("Invalid update")
		return nil
	}

	// Menjalankan query UPDATE untuk mengedit title dan description pada baris yang memiliki id yang sama dengan id yang dikirimkan dalam request.
	if title != "" && description == "" {
		_, err = handlers.DB.Exec("UPDATE TodoList SET title=? WHERE id=?", title, id)
	} else if description != "" && title == "" {
		_, err = handlers.DB.Exec("UPDATE TodoList SET description=? WHERE id=?", description, id)
	} else {
		_, err = handlers.DB.Exec("UPDATE TodoList SET title=?, description=? WHERE id=?", title, description, id)
	}

	// Jika terjadi error saat menjalankan query UPDATE, log error menggunakan logrus
	if err != nil {
		helper.InternalServerError(w, err)
		logrus.Print(err)
	}

	// Prepare response
	apiResponse := model.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    nil,
	}
	logrus.Info("Update Title & Description Todo successfully")

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(apiResponse.Status)
	helper.WriteToResponseBody(w, apiResponse)

	return nil
}

func (handlers *TodoListhandlersImpl) UpdateStatus(w http.ResponseWriter, r *http.Request, config *database.Config) error {

	// Membaca request yang dikirimkan dalam format raw JSON
	var request model.Todos
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		helper.InternalServerError(w, err)
		// Jika terjadi error saat mengambil data, log error tersebut
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error parsing raw JSON request")
		return nil
	}

	// Menyimpan data yang dikirimkan dalam request ke dalam variabel "id" dan "status"
	params := mux.Vars(r)
	id := params["id"]
	status := request.Status

	// Check if Todo exist in the database
	// Menjalankan query untuk menghitung/mencari jumlah baris pada tabel TodoList yang memiliki id yang sama dengan id yang dikirimkan dalam request.
	var count int
	err = handlers.DB.QueryRow("SELECT COUNT(*) FROM TodoList WHERE id=?", id).Scan(&count)
	// Jika terjadi error saat menjalankan query, log error menggunakan logrus dan keluar dari fungsi.
	if err != nil {
		helper.InternalServerError(w, err)
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error checking id")
		return nil
	}

	// Jika tidak ada baris yang memiliki id yang sama dengan id yang dikirimkan dalam request,
	if count == 0 {
		helper.NotFound(w, err)
		return nil
	}
	// Update field Status Todo from the database
	_, err = handlers.DB.Exec("UPDATE TodoList SET status=? WHERE id=?", status, id)

	// Jika terjadi error saat menjalankan query UPDATE, log error menggunakan logrus
	if err != nil {
		helper.InternalServerError(w, err)
		logrus.Print(err)
		return nil
	}
	// Prepare response
	apiResponse := model.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    nil,
	}
	logrus.Info("Update Status Todo successfully")

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(apiResponse.Status)
	helper.WriteToResponseBody(w, apiResponse)

	return nil
}

func (handlers *TodoListhandlersImpl) Delete(w http.ResponseWriter, r *http.Request, config *database.Config) error {

	// Get parameter id from URL
	params := mux.Vars(r)
	id, ok := params["id"]
	// Check if id exists in data
	if !ok {
		helper.BadRequest(w, errors.New(" id not found"))
		logrus.Info(ok)
		return nil
	}

	// Check if Todo exist in the database
	var count int
	if err := handlers.DB.QueryRow("SELECT COUNT(*) FROM TodoList WHERE id=?", id).Scan(&count); err != nil {
		helper.InternalServerError(w, err)
		logrus.WithError(err).Error("Failed to check Todo existence in the database")
		return nil
	}
	// Jika tidak ada baris yang memiliki id yang sama dengan id yang dikirimkan dalam request,
	if count == 0 {
		helper.NotFound(w, errors.New(" id not found in the db"))
		return nil
	}

	// Delete Todo from the database
	if _, err := handlers.DB.Exec("DELETE FROM TodoList WHERE id=?", id); err != nil {
		helper.InternalServerError(w, err)
		logrus.WithError(err).Error("Failed to delete Todo from the database")
		return nil
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

	return nil
}
