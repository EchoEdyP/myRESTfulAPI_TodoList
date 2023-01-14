package handlers

import (
	"RESTfulAPI_todos/error_handling"
	"RESTfulAPI_todos/helper"
	"RESTfulAPI_todos/pkg/database"
	"RESTfulAPI_todos/pkg/model"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

// handler UpdateStatusTodos merupakan fungsi yang menangani request PUT/UPDATE pada API yang ditujukan untuk menghapus todo dari database.
func UpdateStatusTodos(w http.ResponseWriter, r *http.Request, db database.DBConn) {

	conn, err := db.Connect()
	if err != nil {
		error_handling.InternalServerError(w, err)
		logrus.Error(err)
	}
	defer conn.Close()

	// Membaca request yang dikirimkan dalam format raw JSON
	var request model.Todos
	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		error_handling.InternalServerError(w, err)
		// Jika terjadi error saat mengambil data, log error tersebut
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error parsing raw JSON request")
		return
	}

	// Menyimpan data yang dikirimkan dalam request ke dalam variabel "id" dan "status"
	id := request.Id
	status := request.Status

	// Menjalankan query untuk menghitung/mencari jumlah baris pada tabel TodoList yang memiliki id yang sama dengan id yang dikirimkan dalam request.
	var count int
	err = conn.QueryRow("SELECT COUNT(*) FROM TodoList WHERE id=?", id).Scan(&count)
	// Jika terjadi error saat menjalankan query, log error menggunakan logrus dan keluar dari fungsi.
	if err != nil {
		error_handling.InternalServerError(w, err)
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error checking id")
		return
	}

	// Jika tidak ada baris yang memiliki id yang sama dengan id yang dikirimkan dalam request,
	if count == 0 {
		error_handling.NotFound(w, err)
		return
	}

	// Menjalankan query UPDATE untuk mengedit ststus pada baris yang memiliki id yang sama dengan id yang dikirimkan dalam request.
	_, err = conn.Exec("UPDATE TodoList SET status=? WHERE id=?", status, id)

	// Jika terjadi error saat menjalankan query UPDATE, log error menggunakan logrus
	if err != nil {
		error_handling.InternalServerError(w, err)
		logrus.Print(err)
		return
	}

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
}