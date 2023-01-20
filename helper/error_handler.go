package helper

import (
	"RESTfulAPI_todos/pkg/model"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	apiResponse := model.Response{
		Status:  http.StatusInternalServerError,
		Message: "INTERNAL SERVER ERROR",
		Data:    err,
	}
	w.WriteHeader(apiResponse.Status)
	WriteToResponseBody(w, apiResponse)
}

func BadRequest(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	apiResponse := model.Response{
		Status:  http.StatusBadRequest,
		Message: "Bad Request",
		Data:    err,
	}
	w.WriteHeader(apiResponse.Status)
	WriteToResponseBody(w, apiResponse)
}

func NotFound(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	apiResponse := model.Response{
		Status:  http.StatusNotFound,
		Message: "Not Found",
		Data:    err,
	}
	w.WriteHeader(apiResponse.Status)
	WriteToResponseBody(w, apiResponse)
}
