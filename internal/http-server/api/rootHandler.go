package api

import "net/http"

// RootHandler - show is api working
// @Summary This API can be used as health check for this application.
// @Description Tells if the chi-swagger APIs are working or not.
// @Tags info
// @Accept  json
// @Produce  json
// @Success 200 {string} response "api response"
// @Router / [get]
func RootHandler(w http.ResponseWriter, r *http.Request) {
	JSONResponse(w, Response{Success: true, Message: "api is working!"}, http.StatusOK)
}
