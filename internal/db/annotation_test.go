package db_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"papercup-test/internal/model"
)

const videoHost = "www.vimeo.com/"

func TestAnnotationDatabaseBasicCrud(t *testing.T) {

	t.Run("test create", func(t *testing.T) {
		videoID := createVideo(t, videoHost+uuid.NewString()).ID
		createAnnotation(t, videoID)
	})

	t.Run("test get", func(t *testing.T) {
		videoID := createVideo(t, videoHost+uuid.NewString()).ID
		created := createAnnotation(t, videoID)
		returned, err := database.GetAnnotation(created.ID)
		require.NoError(t, err)
		validateAnnotation(t, returned)
	})

	t.Run("test delete", func(t *testing.T) {
		videoID := createVideo(t, videoHost+uuid.NewString()).ID
		created := createAnnotation(t, videoID)
		err := database.DeleteAnnotation(created.ID)
		require.NoError(t, err)
		annotation, err := database.GetAnnotation(created.ID)
		require.Error(t, err)
		require.Nil(t, annotation)
	})

	t.Run("test update", func(t *testing.T) {
		videoID := createVideo(t, videoHost+uuid.NewString()).ID
		created := createAnnotation(t, videoID)
		updates := &model.AnnotationMetadata{
			StartTime: "11:13:12",
			EndTime:   "11:18:10",
			Type:      "commercial",
			Notes:     "Goodbye",
		}
		err := database.UpdateAnnotation(created.ID, updates)
		require.NoError(t, err)
		updated, err := database.GetAnnotation(created.ID)
		require.NoError(t, err)
		validateAnnotationUpdate(t, created, updated, updates)
	})

	t.Run("test get annotation video ID", func(t *testing.T) {
		expectedVideoID := createVideo(t, videoHost+uuid.NewString()).ID
		annotation := createAnnotation(t, expectedVideoID)
		require.NotNil(t, annotation.ID)
		videoID, err := database.GetAnnotationVideoID(annotation.ID)
		require.NoError(t, err)
		require.Equal(t, &expectedVideoID, videoID)
	})

}

func TestAnnotationDatabaseListing(t *testing.T) {
	t.Run("test listing annotations for a video", func(t *testing.T) {
		videoID := createVideo(t, videoHost+uuid.NewString()).ID
		createAnnotationWithTimes(t, "08:08:08", "09:09:09", videoID)
		createAnnotationWithTimes(t, "10:10:10", "11:11:11", videoID)
		createAnnotationWithTimes(t, "11:12:12", "12:12:12", videoID)
		results, err := database.ListAnnotations(videoID)
		require.NoError(t, err)
		require.Equal(t, 3, len(results))
	})
}

func createVideo(t *testing.T, url string) *model.Video {
	t.Helper()
	video, err := database.CreateVideo(&model.Video{
		Name:   "test-video",
		URL:    url,
		Length: "12:12:12",
	})
	require.NoError(t, err)
	require.NotNil(t, video.ID)
	return video
}

func createAnnotationWithTimes(t *testing.T, startTime, endTime string, videoID uuid.UUID) {
	toCreate := &model.Annotation{
		VideoID: videoID,
		AnnotationMetadata: model.AnnotationMetadata{
			StartTime: startTime,
			EndTime:   endTime,
			Type:      "advertisement",
			Notes:     "hello there",
		},
	}
	created, err := database.CreateAnnotation(toCreate)
	require.NoError(t, err)
	validateAnnotation(t, created)
}

func createAnnotation(t *testing.T, videoID uuid.UUID) *model.Annotation {
	toCreate := &model.Annotation{
		VideoID: videoID,
		AnnotationMetadata: model.AnnotationMetadata{
			StartTime: "11:12:12",
			EndTime:   "11:16:10",
			Type:      "advertisement",
			Notes:     "hello there",
		},
	}
	created, err := database.CreateAnnotation(toCreate)
	require.NoError(t, err)
	validateAnnotation(t, created)
	return created
}

func validateAnnotation(t *testing.T, annotation *model.Annotation) {
	t.Helper()
	require.NotNil(t, annotation.ID)
	require.NotNil(t, annotation.VideoID)
	require.NotNil(t, annotation.Type)
	require.NotNil(t, annotation.StartTime)
	require.NotNil(t, annotation.EndTime)
	require.NotNil(t, annotation.Notes)
}

func validateAnnotationUpdate(t *testing.T, original, updated *model.Annotation, updates *model.AnnotationMetadata) {
	t.Helper()
	require.NotEqual(t, original, updated)
	require.Equal(t, original.ID, updated.ID)
	require.Equal(t, original.VideoID, updated.VideoID)
	require.Equal(t, &updated.AnnotationMetadata, updates)
}
