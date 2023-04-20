package auth

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	mockdb "papercup-test/internal/db/mocks"
	"papercup-test/internal/model"
)

func TestAuthoriseUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mockdb.NewMockStore(ctrl)
	service := NewAuthService(m, "test.com", "potato")

	t.Run("should authorise", func(t *testing.T) {
		userID := uuid.New()
		m.EXPECT().GetUserID("mat", "password").Return(&userID, nil)
		creds := &model.Credentials{
			User:     "mat",
			Password: "password",
		}
		token, err := service.AuthoriseUser(creds)
		require.NoError(t, err)
		require.NotEmpty(t, token.Token)
	})

	t.Run("should not authorise", func(t *testing.T) {
		m.EXPECT().GetUserID("mat", "password").Return(nil, errors.New("error"))
		creds := &model.Credentials{
			User:     "mat",
			Password: "password",
		}
		token, err := service.AuthoriseUser(creds)
		require.Error(t, err)
		require.Nil(t, token)
	})

}
