package server

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Conts struct {
	Conts []int
	Need  int
}

func (s *Server) call(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e Conts
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}

	res := s.picker.Calculate(e.Conts, e.Need)

	jsonData, err := json.Marshal(res)
	if err != nil {
		errorResponse(w, "Bad Request "+err.Error(), http.StatusInternalServerError)
		return
	}
	OkResponse(w, jsonData)

}

func OkResponse(w http.ResponseWriter, message []byte) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)

	resp := make(map[string]string)
	resp["Status"] = message
	jsonResp, _ := json.Marshal(resp)

	w.Write(jsonResp)
}
