package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"

	"papercup-test/internal/db"
	mock_db "papercup-test/internal/db/mocks"
	"papercup-test/internal/model"
	"papercup-test/internal/service/auth"
	"papercup-test/internal/service/video"
)

const domain = "test.com"
const secret = "test-secret"

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	server := httptest.NewServer(h)
	return &testServer{server}
}

func newTestApplication(t *testing.T, store db.Store) *app {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	videoService := video.NewService(store)
	authService := auth.NewAuthService(store, domain, secret)
	router := chi.NewRouter()
	logger := log.Logger{}
	app := newApp(router, authService, videoService, &logger, domain, secret)
	return app
}

func (t *testServer) getToken(m *mock_db.MockStore, req *http.Request) {
	userId := uuid.New()
	m.EXPECT().GetUserID("mat", "password").Return(&userId, nil)
	creds := &model.Credentials{
		User:     "mat",
		Password: "password",
	}
	b, _ := json.Marshal(creds)
	reader := bytes.NewReader(b)
	tokenReq, _ := http.NewRequest(http.MethodPost, t.URL+"/authenticate", reader)
	resp, _ := t.Client().Do(tokenReq)
	b, _ = io.ReadAll(resp.Body)
	var token model.Token
	_ = json.Unmarshal(b, &token)
	req.Header.Set("Authorization", "Bearer "+token.Token)
}
