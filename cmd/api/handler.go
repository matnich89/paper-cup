package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"papercup-test/internal/helper"
	"papercup-test/internal/model"
)

var ErrNoResults = errors.New("not found")

func (a *app) Authenticate(w http.ResponseWriter, r *http.Request) {

	var credentials model.Credentials

	err := helper.ReadJSON(w, r, &credentials)
	if err != nil {
		a.unauthorisedResponse(w, r)
		return
	}

	token, err := a.authService.AuthoriseUser(&credentials)
	if err != nil {
		a.unauthorisedResponse(w, r)
		return
	}

	err = helper.WriteJSON(w, http.StatusOK, token)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *app) CreateVideo(w http.ResponseWriter, r *http.Request) {
	var newVideo model.Video
	err := helper.ReadJSON(w, r, &newVideo)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	_, err = url.ParseRequestURI(newVideo.URL)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	createdVideo, err := a.service.CreateVideo(newVideo)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			a.conflictResponse(w, r)
		}
		if strings.HasPrefix(err.Error(), "invalid") {
			a.badRequestResponse(w, r, err)
		}
		a.serverErrorResponse(w, r, err)
		return
	}
	err = helper.WriteJSON(w, http.StatusCreated, createdVideo)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
}

func (a *app) DeleteVideo(w http.ResponseWriter, r *http.Request) {
	videoID, err := uuid.Parse(chi.URLParam(r, "videoID"))
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	err = a.service.DeleteVideo(videoID)
	if err != nil {
		if helper.IsErrorNotFound(err) {
			a.notFoundResponse(w, r)
		} else {
			a.serverErrorResponse(w, r, err)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (a *app) CreateAnnotation(w http.ResponseWriter, r *http.Request) {
	var newAnnotation model.Annotation
	err := helper.ReadJSON(w, r, &newAnnotation)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}
	created, err := a.service.CreateAnnotation(newAnnotation)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			a.badRequestResponse(w, r, errors.New("invalid video id"))
			return
		}
		if strings.HasPrefix(err.Error(), "invalid") {
			a.badRequestResponse(w, r, err)
		}
		if strings.Contains(err.Error(), "duplicate") {
			a.conflictResponse(w, r)
			return
		}
		a.serverErrorResponse(w, r, err)
	}
	err = helper.WriteJSON(w, http.StatusCreated, created)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *app) UpdateAnnotation(w http.ResponseWriter, r *http.Request) {
	annotationID, err := uuid.Parse(chi.URLParam(r, "annotationID"))
	var updates model.AnnotationMetadata
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}
	err = a.service.UpdateAnnotation(annotationID, updates)
	if err != nil {
		if err.Error() == ErrNoResults.Error() {
			a.notFoundResponse(w, r)
			return
		}
		if strings.HasPrefix(err.Error(), "invalid") {
			a.badRequestResponse(w, r, err)
			return
		}
		a.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *app) DeleteAnnotation(w http.ResponseWriter, r *http.Request) {
	annotationID, err := uuid.Parse(chi.URLParam(r, "annotationID"))
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	err = a.service.DeleteAnnotation(annotationID)

	if err != nil {
		if helper.IsErrorNotFound(err) {
			a.notFoundResponse(w, r)
		} else {
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (a *app) ListAnnotation(w http.ResponseWriter, r *http.Request) {
	videoID, err := uuid.Parse(chi.URLParam(r, "videoID"))
	if err != nil {
		a.badRequestResponse(w, r, err)
	}

	annotations, err := a.service.ListAnnotations(videoID)
	if err != nil {
		if err.Error() == ErrNoResults.Error() {
			a.notFoundResponse(w, r)
			return
		}
		a.serverErrorResponse(w, r, err)
	}
	err = helper.WriteJSON(w, http.StatusOK, annotations)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
}
