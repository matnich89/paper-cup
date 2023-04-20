package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/pascaldekloe/jwt"
)

func (a *app) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Authorization")
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func (a *app) checkToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader == "" {
			a.unauthorisedResponse(w, r)
			return
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			a.unauthorisedResponse(w, r)
			return
		}

		token := headerParts[1]

		claims, err := jwt.HMACCheck([]byte(token), []byte(a.secret))
		if err != nil {
			a.unauthorisedResponse(w, r)
			return
		}

		if !claims.Valid(time.Now()) {
			a.unauthorisedResponse(w, r)
			return
		}

		if claims.Issuer != a.domain {
			a.unauthorisedResponse(w, r)
			return
		}

		if !claims.AcceptAudience(a.domain) {
			a.unauthorisedResponse(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
