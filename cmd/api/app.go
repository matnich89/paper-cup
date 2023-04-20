package main

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"papercup-test/internal/service/auth"
	"papercup-test/internal/service/video"
)

type app struct {
	router      *chi.Mux
	authService *auth.Service
	service     *video.Service
	logger      *log.Logger
	domain      string
	secret      string
}

func newApp(router *chi.Mux, authService *auth.Service,
	service *video.Service, logger *log.Logger, domain, secret string) *app {
	return &app{router: router, authService: authService, service: service, logger: logger, domain: domain, secret: secret}
}

func (a *app) routes() *chi.Mux {

	a.router.Use(middleware.Recoverer)
	a.router.Use(a.enableCORS)
	a.router.Group(func(r chi.Router) {
		r.Post("/authenticate", a.Authenticate)
	})
	a.router.Group(func(r chi.Router) {
		r.Use(a.checkToken)
		r.Post("/videos", a.CreateVideo)
		r.Delete("/videos/{videoID}", a.DeleteVideo)
	})

	a.router.Group(func(r chi.Router) {
		r.Use(a.checkToken)
		r.Post("/annotations", a.CreateAnnotation)
		r.Get("/annotations/{videoID}", a.ListAnnotation)
		r.Put("/annotations/{annotationID}", a.UpdateAnnotation)
		r.Delete("/annotations/{annotationID}", a.DeleteAnnotation)
	})

	return a.router
}
