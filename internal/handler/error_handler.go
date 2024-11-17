package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Temutjin2k/doodocs_Challange/models"
)

func SendError(w http.ResponseWriter, ErrorText string, code int) {
	Error := models.Error{
		Message: ErrorText,
	}
	jsondata, _ := json.MarshalIndent(Error, "", "    ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsondata)
}
