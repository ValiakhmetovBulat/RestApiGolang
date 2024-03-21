package api

import (
	log "RestApiGolang/internal/logger"
	"RestApiGolang/internal/models"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/jinzhu/gorm"
	"net/http"
)

// Redirect redirects user by given alias
// @Summary Redirect
// @Tags redirect
// @Description redirects user by given alias
// @ID redirect-user-by-alias
// @Accept json
// @Produce json
// @Param alias path string true "alias"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /{alias} [get]
func Redirect(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http-server.api.Redirect"
	alias := chi.URLParam(r, "alias")
	if alias == "" {
		JSONResponse(w, Response{Success: false, Message: ErrAliasIsEmpty}, http.StatusBadRequest)
		log.Errorf("%s: %s", op, ErrAliasIsEmpty)
		return
	}

	url, err := models.GetURL(alias)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		JSONResponse(w, Response{Success: false, Message: ErrUrlNotFound}, http.StatusNotFound)
		log.Errorf("%s: %s", op, err.Error())
		return
	}
	if err != nil {
		JSONResponse(w, Response{Success: false, Message: ErrInternalError}, http.StatusInternalServerError)
		log.Errorf("%s: %s", op, err.Error())
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}
