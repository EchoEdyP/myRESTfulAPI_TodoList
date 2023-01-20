package handlers

import (
	"RESTfulAPI_todos/helper"
	"RESTfulAPI_todos/pkg/database"
	"RESTfulAPI_todos/pkg/model"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

// handler UpdateStatusTodos merupakan fungsi yang menangani request PUT/UPDATE pada API yang ditujukan untuk menghapus todo dari database.
func UpdateTitleDescriptionTodos(w http.ResponseWriter, r *http.Request, config *database.Config) {

	conn, err := database.ConnectDB(config)
	if err != nil {
		helper.InternalServerError(w, err)
		logrus.Error(err)
	}
	defer conn.Close()

	// Membaca request yang dikirimkan dalam format raw JSON
	var request model.Todos
	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		helper.InternalServerError(w, err)
		// Jika terjadi error saat mengambil data, log error tersebut
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error parsing raw JSON request")
		return
	}

	// Menyimpan data yang dibaca dari request ke dalam variabel id, title, dan description.
	params := mux.Vars(r)
	id := params["id"]
	title := request.Title
	description := request.Description

	// Check if Todo exist in the database
	// Menjalankan query untuk menghitung/mencari jumlah baris pada tabel TodoList yang memiliki id yang sama dengan id yang dikirimkan dalam request.
	var count int
	err = conn.QueryRow("SELECT COUNT(*) FROM TodoList WHERE id=?", id).Scan(&count)
	// Jika terjadi error saat menjalankan query, log error menggunakan logrus dan keluar dari fungsi.
	if err != nil {
		helper.InternalServerError(w, err)
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error checking id")
		return
	}

	// Jika tidak ada baris yang memiliki id yang sama dengan id yang dikirimkan dalam request,
	if count == 0 {
		helper.NotFound(w, err)
		return
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
		return
	}

	// Menjalankan query UPDATE untuk mengedit title dan description pada baris yang memiliki id yang sama dengan id yang dikirimkan dalam request.
	if title != "" && description == "" {
		_, err = conn.Exec("UPDATE TodoList SET title=? WHERE id=?", title, id)
	} else if description != "" && title == "" {
		_, err = conn.Exec("UPDATE TodoList SET description=? WHERE id=?", description, id)
	} else {
		_, err = conn.Exec("UPDATE TodoList SET title=?, description=? WHERE id=?", title, description, id)
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
}
