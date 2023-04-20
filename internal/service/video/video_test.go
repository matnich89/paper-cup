package video

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	mockdb "papercup-test/internal/db/mocks"
	"papercup-test/internal/model"
)

func TestCreateVideo(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mockdb.NewMockStore(ctrl)
	service := NewService(m)

	t.Run("should return created video", func(t *testing.T) {
		video := createValidVideo()
		m.EXPECT().CreateVideo(video).Return(video, nil)

		actual, err := service.CreateVideo(*video)
		require.NoError(t, err)
		require.NotNil(t, actual)
	})

	t.Run("should handle error when creating video", func(t *testing.T) {
		video := createValidVideo()
		m.EXPECT().CreateVideo(video).Return(nil, errors.New("error"))

		actual, err := service.CreateVideo(*video)
		require.Nil(t, actual)
		require.Error(t, err)
	})

	t.Run("should handle invalid length hour", func(t *testing.T) {
		video := createVideoWithLength("-40:01:01")
		_, err := service.CreateVideo(*video)
		require.Equal(t, "invalid hours provided", err.Error())
	})

	t.Run("should handle invalid length minutes", func(t *testing.T) {
		video := createVideoWithLength("12:64:00")
		_, err := service.CreateVideo(*video)
		require.Equal(t, "invalid minutes provided", err.Error())
	})

	t.Run("should handle invalid length seconds", func(t *testing.T) {
		video := createVideoWithLength("12:12:78")
		_, err := service.CreateVideo(*video)
		require.Equal(t, "invalid seconds provided", err.Error())
	})

}

func TestDeleteVideo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mockdb.NewMockStore(ctrl)
	service := NewService(m)

	t.Run("should handle successful deletion of video", func(t *testing.T) {
		videoID := uuid.New()
		m.EXPECT().DeleteVideo(videoID).Return(nil)

		err := service.DeleteVideo(videoID)
		require.NoError(t, err)
	})

	t.Run("should handle error when deleting video", func(t *testing.T) {
		videoID := uuid.New()
		m.EXPECT().DeleteVideo(videoID).Return(errors.New("error"))
		err := service.DeleteVideo(videoID)
		require.Error(t, err)
	})
}

func TestCreateAnnotation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mockdb.NewMockStore(ctrl)
	service := NewService(m)

	t.Run("should return created annotation", func(t *testing.T) {
		videoID := uuid.New()
		video := createValidVideoWithId(videoID)
		m.EXPECT().GetVideoLength(videoID).Return(video.Length, nil)
		annotation := createAnnotation("11:11:11", "12:12:12", videoID)
		m.EXPECT().CreateAnnotation(annotation).Return(annotation, nil)
		created, err := service.CreateAnnotation(*annotation)
		require.NoError(t, err)
		require.NotNil(t, created)
	})

	t.Run("should handle out of bounds annotation start time", func(t *testing.T) {
		videoID := uuid.New()
		video := createValidVideoWithId(videoID)
		m.EXPECT().GetVideoLength(videoID).Return(video.Length, nil)
		annotation := createAnnotation("13:00:00", "12:12:12", videoID)
		created, err := service.CreateAnnotation(*annotation)
		require.Error(t, err)
		require.Equal(t, "invalid start time provided", err.Error())
		require.Nil(t, created)
	})

	t.Run("should handle out of bounds annotation end time", func(t *testing.T) {
		videoID := uuid.New()
		video := createValidVideoWithId(videoID)
		m.EXPECT().GetVideoLength(videoID).Return(video.Length, nil)
		annotation := createAnnotation("12:00:00", "13:00:00", videoID)
		created, err := service.CreateAnnotation(*annotation)
		require.Error(t, err)
		require.Equal(t, "invalid end time provided", err.Error())
		require.Nil(t, created)
	})

	t.Run("should handle invalid videoID", func(t *testing.T) {
		m.EXPECT().GetVideoLength(gomock.Any()).Return("", errors.New("not found"))
		annotation := createAnnotation("12:00:00", "13:00:00", uuid.New())
		created, err := service.CreateAnnotation(*annotation)
		require.Error(t, err)
		require.Nil(t, created)
	})

}

func TestGetAnnotation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mockdb.NewMockStore(ctrl)
	service := NewService(m)

	t.Run("should get annotation", func(t *testing.T) {
		annotation := createAnnotation("11:11:11", "12:12:12", uuid.New())
		m.EXPECT().GetAnnotation(annotation.ID).Return(annotation, nil)
		result, err := service.GetAnnotation(annotation.ID)
		require.NoError(t, err)
		require.NotNil(t, result)
	})

	t.Run("should handle get annotation error", func(t *testing.T) {
		id := uuid.New()
		m.EXPECT().GetAnnotation(id).Return(nil, errors.New("error"))
		result, err := service.GetAnnotation(id)
		require.Error(t, err)
		require.Nil(t, result)
	})
}

func TestUpdateAnnotation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mockdb.NewMockStore(ctrl)
	service := NewService(m)

	t.Run("should update annotation", func(t *testing.T) {
		videoID := uuid.New()
		video := createValidVideoWithId(videoID)
		m.EXPECT().GetVideoLength(videoID).Return(video.Length, nil)
		update := createValidAnnotationMetadata()
		annotationID := uuid.New()
		m.EXPECT().GetAnnotationVideoID(annotationID).Return(&videoID, nil)
		m.EXPECT().UpdateAnnotation(annotationID, update).Return(nil)
		err := service.UpdateAnnotation(annotationID, *update)
		require.NoError(t, err)
	})

	t.Run("should handle out of bounds annotation update start time", func(t *testing.T) {
		videoID := uuid.New()
		video := createValidVideoWithId(videoID)
		m.EXPECT().GetVideoLength(videoID).Return(video.Length, nil)
		update := createAnnotationMetaDataWithTimes("13:00:00", "12:00:00")
		annotationID := uuid.New()
		m.EXPECT().GetAnnotationVideoID(annotationID).Return(&videoID, nil)
		err := service.UpdateAnnotation(annotationID, *update)
		require.Error(t, err)
	})

	t.Run("should handle out of bounds annotation update end time", func(t *testing.T) {
		videoID := uuid.New()
		video := createValidVideoWithId(videoID)
		m.EXPECT().GetVideoLength(videoID).Return(video.Length, nil)
		update := createAnnotationMetaDataWithTimes("12:00:00", "13:00:00")
		annotationID := uuid.New()
		m.EXPECT().GetAnnotationVideoID(annotationID).Return(&videoID, nil)
		err := service.UpdateAnnotation(annotationID, *update)
		require.Error(t, err)
	})
}

func TestListAnnotations(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mockdb.NewMockStore(ctrl)
	service := NewService(m)

	t.Run("should list annotations for video", func(t *testing.T) {
		videoID := uuid.New()
		annotation1 := createAnnotation("10:00:00", "10:30:00", videoID)
		annotation2 := createAnnotation("11:11:11", "12:12:12", videoID)

		m.EXPECT().ListAnnotations(videoID).Return([]*model.Annotation{annotation1, annotation2}, nil)
		result, err := service.ListAnnotations(videoID)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Len(t, result, 2)
	})

	t.Run("should handle error when listing annotations", func(t *testing.T) {
		videoID := uuid.New()
		m.EXPECT().ListAnnotations(videoID).Return(nil, errors.New("error"))
		result, err := service.ListAnnotations(videoID)
		require.Nil(t, result)
		require.Error(t, err)
	})

}

func createAnnotation(startTime, endTime string, videoID uuid.UUID) *model.Annotation {
	return &model.Annotation{
		ID:      uuid.New(),
		VideoID: videoID,
		AnnotationMetadata: model.AnnotationMetadata{
			StartTime: startTime,
			EndTime:   endTime,
			Type:      "advertisement",
			Notes:     "Hello There",
		},
	}
}

func createAnnotationMetaDataWithTimes(startTime, endTime string) *model.AnnotationMetadata {
	return &model.AnnotationMetadata{
		StartTime: startTime,
		EndTime:   endTime,
		Type:      "commercial",
		Notes:     "Goodbye",
	}
}

func createValidAnnotationMetadata() *model.AnnotationMetadata {
	return &model.AnnotationMetadata{
		StartTime: "12:00:00",
		EndTime:   "12:00:10",
		Type:      "commercial",
		Notes:     "Goodbye",
	}
}

func createValidVideo() *model.Video {
	return &model.Video{
		ID:     uuid.New(),
		Name:   "test-video",
		URL:    "http://www.youtube.com/123456",
		Length: "12:12:12",
	}
}

func createValidVideoWithId(id uuid.UUID) *model.Video {
	return &model.Video{
		ID:     id,
		Name:   "test-video",
		URL:    "http://www.youtube.com/123456",
		Length: "12:12:12",
	}
}

func createVideoWithLength(length string) *model.Video {
	return &model.Video{
		ID:     uuid.New(),
		Name:   "test-video",
		URL:    "http://www.youtube.com/123456",
		Length: length,
	}
}
