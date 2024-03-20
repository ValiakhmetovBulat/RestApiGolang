package api

import "net/http"

func RootHandler(w http.ResponseWriter, r *http.Request) {
	JSONResponse(w, Response{Success: true, Message: "api is working!"}, http.StatusOK)
	return
}
