package api

import (
	"encoding/json"
	"net/http"
)

func writeResp(w http.ResponseWriter, status int, msg any, key string) {
	env := map[string]any{key: msg}
	bs, err := json.Marshal(env)
	if err != nil {
		http.Error(w, "error marshaling the response into json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(bs)
	if err != nil {
		http.Error(w, "error marshaling the response into json", http.StatusInternalServerError)
		return
	}
}

func writeError(w http.ResponseWriter, status int, errorMsg string) {
	writeResp(w, status, errorMsg, "error")
}

func writeServerError(w http.ResponseWriter) {
	writeError(w, http.StatusInternalServerError, "internal server error")
}

func writeBadRequestError(w http.ResponseWriter) {
	writeError(w, http.StatusBadRequest, "bad request")
}
