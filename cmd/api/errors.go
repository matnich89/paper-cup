package main

import (
	"net/http"

	"papercup-test/internal/helper"
)

type Wrapper struct {
	Error string `json:"error"`
}

func (a *app) errorResponse(w http.ResponseWriter, r *http.Request, status int, message string) {
	wrapper := &Wrapper{Error: message}
	err := helper.WriteJSON(w, status, wrapper)
	if err != nil {
		a.logger.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (a *app) unauthorisedResponse(w http.ResponseWriter, r *http.Request) {
	message := "you must be authenticated to access this resource"
	a.errorResponse(w, r, http.StatusUnauthorized, message)
}

func (a *app) conflictResponse(w http.ResponseWriter, r *http.Request) {
	message := "there is a conflict with another record"
	a.errorResponse(w,r, http.StatusConflict, message)
}

func (a *app) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	a.logger.Println("internal server error ", err.Error())
	message := "the server encountered a problem and could not process your request"
	a.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (a *app) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	a.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (a *app) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	a.errorResponse(w, r, http.StatusNotFound, message)
}
