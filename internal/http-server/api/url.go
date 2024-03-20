package api

import (
	"RestApiGolang/internal/models"
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"net/http"
)

type RequestBody struct {
	Alias string `json:"alias"`
}

const (
	ErrUrlNotFound   string = "url not found"
	ErrInternalError string = "server error"
	ErrInvalidJSON   string = "invalid JSON structure"
)

// GetUrlByAlias gets url by given alias
// @Summary GetUrlByAlias
// @Tags url
// @Description gets url by alias
// @ID get-url-by-alias
// @Accept json
// @Produce json
// @Param input body RequestBody true "alias"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /url [get]
func GetUrlByAlias(w http.ResponseWriter, r *http.Request) {
	var rb RequestBody

	err := json.NewDecoder(r.Body).Decode(&rb)
	if err != nil {
		JSONResponse(w, Response{Success: false, Message: ErrInvalidJSON}, http.StatusBadRequest)
		return
	}

	u, err := models.GetURL(rb.Alias)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		JSONResponse(w, Response{Success: false, Message: ErrUrlNotFound}, http.StatusNotFound)
		return
	}
	if err != nil {
		JSONResponse(w, Response{Success: false, Message: ErrInternalError}, http.StatusInternalServerError)
		return
	}
	JSONResponse(w, u, http.StatusOK)
	return
}

// DeleteUrlByAlias deletes url by given alias
// @Summary DeleteUrlByAlias
// @Tags url
// @Description deletes url by alias
// @ID delete-url-by-alias
// @Accept json
// @Produce json
// @Param input body RequestBody true "alias"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /url [delete]
func DeleteUrlByAlias(w http.ResponseWriter, r *http.Request) {
	var rb RequestBody

	err := json.NewDecoder(r.Body).Decode(&rb)
	if err != nil {
		JSONResponse(w, Response{Success: false, Message: ErrInvalidJSON}, http.StatusBadRequest)
		return
	}

	err = models.DeleteURL(rb.Alias)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		JSONResponse(w, Response{Success: false, Message: ErrUrlNotFound}, http.StatusInternalServerError)
		return
	}
	if err != nil {
		JSONResponse(w, Response{Success: false, Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	JSONResponse(w, Response{Success: true, Message: "url deleted!"}, http.StatusOK)
}

// PostUrl creates url by given JSON
// @Summary PostUrl
// @Tags urls
// @Description posts url
// @ID post-url
// @Accept json
// @Produce json
// @Param input body models.Url true "url"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /urls [post]
func PostUrl(w http.ResponseWriter, r *http.Request) {
	u := models.Url{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		JSONResponse(w, Response{Success: false, Message: ErrInvalidJSON}, http.StatusBadRequest)
		return
	}
	err = models.SaveURL(&u)
	if errors.Is(err, models.ErrAliasNotSpecified) || errors.Is(err, models.ErrUrlNotSpecified) {
		JSONResponse(w, Response{Success: false, Message: err.Error()}, http.StatusBadRequest)
		return
	}
	if err != nil {
		JSONResponse(w, Response{Success: false, Message: "error inserting url into database"}, http.StatusInternalServerError)
		return
	}
	JSONResponse(w, u, http.StatusCreated)
}

// PutUrl updates url by given JSON
// @Summary PutUrl
// @Tags url
// @Description updates url
// @ID put-url
// @Accept json
// @Produce json
// @Param input body models.Url true "url"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /url [put]
func PutUrl(w http.ResponseWriter, r *http.Request) {
	u := models.Url{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		JSONResponse(w, Response{Success: false, Message: ErrInvalidJSON}, http.StatusBadRequest)
		return
	}
	err = models.PutURL(&u)
	if err != nil {
		JSONResponse(w, Response{Success: false, Message: ErrInternalError}, http.StatusInternalServerError)
		return
	}
	JSONResponse(w, u, http.StatusOK)

}

// GetUrls gets all urls
// @Summary GetUrls
// @Tags urls
// @Description gets all urls
// @ID get-all-urls
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /urls [get]
func GetUrls(w http.ResponseWriter, r *http.Request) {
	us, err := models.GetURLs()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		JSONResponse(w, Response{Success: false, Message: ErrUrlNotFound}, http.StatusNotFound)
		return
	}
	if err != nil {
		JSONResponse(w, Response{Success: false, Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	JSONResponse(w, us, http.StatusOK)
	return
}

func URLByAlias(w http.ResponseWriter, r *http.Request) {
	var rb RequestBody

	err := json.NewDecoder(r.Body).Decode(&rb)
	if err != nil {
		JSONResponse(w, Response{Success: false, Message: ErrInvalidJSON}, http.StatusBadRequest)
		return
	}
	switch {
	case r.Method == "GET":
		u, err := models.GetURL(rb.Alias)
		if err != nil {
			JSONResponse(w, Response{Success: false, Message: ErrInternalError}, http.StatusInternalServerError)
			return
		}
		JSONResponse(w, u, http.StatusOK)
		return
	case r.Method == "DELETE":
		err = models.DeleteURL(rb.Alias)
		if err != nil {
			JSONResponse(w, Response{Success: false, Message: ErrInternalError}, http.StatusInternalServerError)
			return
		}
		JSONResponse(w, Response{Success: true, Message: "url deleted!"}, http.StatusOK)
	}
}

func URL(w http.ResponseWriter, r *http.Request) {
	u := models.Url{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		JSONResponse(w, Response{Success: false, Message: ErrInvalidJSON}, http.StatusBadRequest)
		return
	}
	switch {
	case r.Method == "PUT":
		err = models.PutURL(&u)
		if err != nil {
			JSONResponse(w, Response{Success: false, Message: err.Error()}, http.StatusInternalServerError)
			return
		}
		JSONResponse(w, u, http.StatusOK)
	}
}

func URLs(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		us, err := models.GetURLs()
		if err != nil {
			JSONResponse(w, Response{Success: false, Message: err.Error()}, http.StatusInternalServerError)
			return
		}
		JSONResponse(w, us, http.StatusOK)
		return
	case r.Method == "POST":
		u := models.Url{}
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			JSONResponse(w, Response{Success: false, Message: "invalid JSON structure"}, http.StatusBadRequest)
			return
		}
		err = models.SaveURL(&u)
		if err == models.ErrAliasNotSpecified || err == models.ErrUrlNotSpecified {
			JSONResponse(w, Response{Success: false, Message: err.Error()}, http.StatusBadRequest)
			return
		}
		if err != nil {
			JSONResponse(w, Response{Success: false, Message: "error inserting url into database"}, http.StatusInternalServerError)
			return
		}
		JSONResponse(w, u, http.StatusCreated)
	}
}
