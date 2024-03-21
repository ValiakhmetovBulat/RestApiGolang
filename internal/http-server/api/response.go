package api

import (
	log "RestApiGolang/internal/logger"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

const (
	ErrUrlNotFound   string = "url not found"
	ErrInternalError string = "server error"
	ErrInvalidJSON   string = "invalid JSON structure"
	ErrAliasIsEmpty  string = "alias is empty"
)

func JSONResponse(w http.ResponseWriter, d interface{}, c int) {
	dj, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		http.Error(w, "error creating JSON response", http.StatusInternalServerError)
		log.Error(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj)
}
