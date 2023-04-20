package video

import (
	"errors"
	"strconv"
	"strings"

	"github.com/google/uuid"

	"papercup-test/internal/db"
	"papercup-test/internal/model"
)

type Service struct {
	store db.Store
}

func NewService(store db.Store) *Service {
	return &Service{store: store}
}

func (s *Service) CreateVideo(video model.Video) (*model.Video, error) {
	err := validateLength(video.Length)
	if err != nil {
		return nil, err
	}
	createdVideo, err := s.store.CreateVideo(&video)
	if err != nil {
		return nil, err
	}
	return createdVideo, nil
}

func (s *Service) DeleteVideo(videoID uuid.UUID) error {
	err := s.store.DeleteVideo(videoID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) CreateAnnotation(annotation model.Annotation) (*model.Annotation, error) {
	err := validateLength(annotation.StartTime)
	if err != nil {
		return nil, err
	}

	err = validateLength(annotation.EndTime)
	if err != nil {
		return nil, err
	}

	err = s.validateAnnotationTimes(annotation.StartTime, annotation.EndTime, annotation.VideoID)
	if err != nil {
		return nil, err
	}

	err = validateEndTimeAfterStartTime(annotation.StartTime, annotation.EndTime)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	created, err := s.store.CreateAnnotation(&annotation)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (s *Service) GetAnnotation(id uuid.UUID) (*model.Annotation, error) {
	annotation, err := s.store.GetAnnotation(id)
	if err != nil {
		return nil, err
	}
	return annotation, nil
}

func (s *Service) ListAnnotations(videoID uuid.UUID) ([]*model.Annotation, error) {
	results, err := s.store.ListAnnotations(videoID)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (s *Service) UpdateAnnotation(id uuid.UUID, updates model.AnnotationMetadata) error {
	err := validateLength(updates.StartTime)
	if err != nil {
		return err
	}

	err = validateLength(updates.EndTime)
	if err != nil {
		return err
	}

	err = validateEndTimeAfterStartTime(updates.StartTime, updates.EndTime)
	if err != nil {
		return err
	}

	videoID, err := s.store.GetAnnotationVideoID(id)
	if err != nil {
		return err
	}

	err = s.validateAnnotationTimes(updates.StartTime, updates.EndTime, *videoID)
	if err != nil {
		return err
	}

	err = s.store.UpdateAnnotation(id, &updates)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteAnnotation(id uuid.UUID) error {
	err := s.store.DeleteAnnotation(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) validateAnnotationTimes(startTime, endTime string, videoID uuid.UUID) error {
	videoLength, err := s.store.GetVideoLength(videoID)
	if err != nil {
		return err
	}

	videoLengthValue := getTimeNumbers(videoLength)
	startTimeValue := getTimeNumbers(startTime)
	if startTimeValue > videoLengthValue {
		return errors.New("invalid start time provided")
	}

	endTimeValue := getTimeNumbers(endTime)
	if endTimeValue > videoLengthValue {
		return errors.New("invalid end time provided")
	}

	return nil
}

func getTimeNumbers(time string) int {
	timeArray := strings.Split(time, ":")
	var numberStr string
	for _, value := range timeArray {
		numberStr = numberStr + value
	}
	number, _ := strconv.Atoi(numberStr)
	return number
}

func validateLength(length string) error {
	lengths := strings.Split(length, ":")
	if len(lengths) != 3 {
		return errors.New("invalid length provided")
	}

	lengthHours, err := strconv.Atoi(lengths[0])

	if err != nil || lengthHours < 0 {
		return errors.New("invalid hours provided")
	}

	lengthMinutes, err := strconv.Atoi(lengths[1])

	if err != nil || lengthMinutes < 0 || lengthMinutes > 59 {
		return errors.New("invalid minutes provided")
	}

	lengthSeconds, err := strconv.Atoi(lengths[2])

	if err != nil || lengthSeconds < 0 || lengthSeconds > 59 {
		return errors.New("invalid seconds provided")
	}
	return nil
}

func validateEndTimeAfterStartTime(startTime, endTime string) error {
	start := getTimeNumbers(startTime)
	end := getTimeNumbers(endTime)
	if end > start {
		return errors.New("invalid end is after start")
	}
	return nil
}
