package helper

import (
	"encoding/json"
	"net/http"
)

func WriteToResponseBody(w http.ResponseWriter, response interface{}) {

	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicIfError(err)
}
