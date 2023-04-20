package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	mock_db "papercup-test/internal/db/mocks"
	"papercup-test/internal/model"
)

const invalidBody = "{\"name\": \"bob\"}"
const videoHost = "http://vimeo.com"

func TestAuthentication(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_db.NewMockStore(ctrl)
	app := newTestApplication(t, m)
	server := newTestServer(t, app.routes())
	defer server.Close()

	t.Run("should authenticate", func(t *testing.T) {
		userId := uuid.New()
		m.EXPECT().GetUserID("mat", "password").Return(&userId, nil)
		creds := &model.Credentials{
			User:     "mat",
			Password: "password",
		}
		b, _ := json.Marshal(creds)
		reader := bytes.NewReader(b)
		req, _ := http.NewRequest(http.MethodPost, server.URL+"/authenticate", reader)
		resp, err := server.Client().Do(req)
		require.NoError(t, err)
		b, err = io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.NotEmpty(t, string(b))
	})

	t.Run("should not authenticate", func(t *testing.T) {
		m.EXPECT().GetUserID(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
		creds := &model.Credentials{
			User:     "mat",
			Password: "password1",
		}
		b, _ := json.Marshal(creds)
		reader := bytes.NewReader(b)
		req, _ := http.NewRequest(http.MethodPost, server.URL+"/authenticate", reader)
		resp, err := server.Client().Do(req)
		require.NoError(t, err)
		require.Equal(t, resp.StatusCode, http.StatusUnauthorized)
	})
}

func TestVideoCreation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_db.NewMockStore(ctrl)
	app := newTestApplication(t, m)
	server := newTestServer(t, app.routes())
	defer server.Close()

	url := server.URL + "/videos"

	var tests = []struct {
		name                    string
		expectedResponse        int
		shouldCheckResponseBody bool
		validPayload            bool
	}{
		{"creation", http.StatusCreated, true, true},
		{"invalid payload", http.StatusBadRequest, false, false},
		{"should handle server error", http.StatusInternalServerError, false, true},
	}

	for _, entry := range tests {
		t.Run(entry.name, func(t *testing.T) {
			video := createVideo(videoHost + uuid.NewString())

			var req *http.Request
			var err error
			if entry.validPayload {
				if entry.expectedResponse == http.StatusCreated {
					m.EXPECT().CreateVideo(video).Return(video, nil)
				} else {
					m.EXPECT().CreateVideo(video).Return(nil, errors.New("err"))
				}
				req, err = createRequest(url, http.MethodPost, video)
			} else {
				req, err = createRequest(url, http.MethodPost, invalidBody)
			}

			require.NoError(t, err)
			server.getToken(m, req)
			resp, err := server.Client().Do(req)
			require.NoError(t, err)
			require.Equal(t, entry.expectedResponse, resp.StatusCode)

			if entry.shouldCheckResponseBody {
				var created = model.Video{}
				err = readResponseBody(resp, &created)
				require.NoError(t, err)
				require.NotNil(t, created)
				require.NotNil(t, created.ID)
			}
		})
	}
}

func TestVideoDeletion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_db.NewMockStore(ctrl)
	app := newTestApplication(t, m)
	server := newTestServer(t, app.routes())
	defer server.Close()

	url := server.URL + "/videos"

	var tests = []struct {
		name              string
		expectedResponse  int
		shouldCreateVideo bool
	}{
		{"deletion", http.StatusNoContent, true},
		{"not found", http.StatusNotFound, false},
	}

	for _, entry := range tests {
		videoID := uuid.New()
		if entry.shouldCreateVideo {
			video := createVideo(videoHost + uuid.NewString())
			req, err := createRequest(url, http.MethodPost, video)
			require.NoError(t, err)

			server.getToken(m, req)
			m.EXPECT().CreateVideo(video).Return(video, nil)

			resp, err := server.Client().Do(req)
			require.NoError(t, err)

			var createdVideo model.Video
			err = readResponseBody(resp, &createdVideo)
			require.NoError(t, err)

			videoID = createdVideo.ID
			m.EXPECT().DeleteVideo(videoID).Return(nil)
		}
		url = url + "/" + videoID.String()
		req, err := createRequest(url, http.MethodDelete, nil)
		server.getToken(m, req)
		require.NoError(t, err)
		resp, err := server.Client().Do(req)
		require.NoError(t, err)
		require.Equal(t, entry.expectedResponse, resp.StatusCode)
	}

}

func createVideo(url string) *model.Video {
	return &model.Video{
		ID:     uuid.New(),
		Name:   "test-video",
		URL:    url,
		Length: "12:12:12",
	}
}

func createRequest(url, method string, body any) (*http.Request, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(b)

	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func readResponseBody(resp *http.Response, data any) error {
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, data)
	if err != nil {
		return err
	}
	return nil
}
