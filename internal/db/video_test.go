package db_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestVideoDatabaseBasicCrud(t *testing.T) {

	t.Run("test create", func(t *testing.T) {
		video := createVideo(t, "www.vimeo.com/" + uuid.NewString())
		require.NotNil(t, video)
		require.NotNil(t, video.ID)
	})

	t.Run("test delete", func(t *testing.T) {
		video := createVideo(t, "www.vimeo.com/" + uuid.NewString())
		err := database.DeleteVideo(video.ID)
		require.NoError(t, err)
	})

	t.Run("test delete none existent row", func(t *testing.T) {
		err := database.DeleteVideo(uuid.New())
		require.Error(t, err)
	})

	t.Run("test get video length", func(t *testing.T) {
		video := createVideo(t, "www.vimeo.com/" + uuid.NewString())
		length, err := database.GetVideoLength(video.ID)
		require.NoError(t, err)
		require.Equal(t, video.Length, length)
	})
}

