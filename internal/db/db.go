package db

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"papercup-test/internal/model"
)

type Store interface {
	CreateVideo(newVideo *model.Video) (*model.Video, error)
	GetVideoLength(videoID uuid.UUID) (string, error)
	GetAnnotationVideoID(annotationID uuid.UUID) (*uuid.UUID, error)
	DeleteVideo(videoID uuid.UUID) error
	CreateAnnotation(newAnnotaton *model.Annotation) (*model.Annotation, error)
	GetAnnotation(id uuid.UUID) (*model.Annotation, error)
	ListAnnotations(videoID uuid.UUID) ([]*model.Annotation, error)
	UpdateAnnotation(annotationID uuid.UUID, updatedAnno *model.AnnotationMetadata) error
	DeleteAnnotation(annotationID uuid.UUID) error
	GetUserID(username, password string) (*uuid.UUID, error)
}

type Database struct {
	Client *sql.DB
}

func NewDatabase(host, port, username, password, sslEnabled string) (*Database, error) {

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s",
		host,
		port,
		username,
		password,
		sslEnabled)

	dbConn, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return &Database{Client: dbConn}, nil
}
