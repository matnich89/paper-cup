package db

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"papercup-test/internal/model"
)

func (d *Database) GetAnnotation(id uuid.UUID) (*model.Annotation, error) {
	query := `SELECT * FROM annotations WHERE annotations.id = $1`

	var annotation model.Annotation

	err := d.Client.QueryRow(query, id).Scan(&annotation.ID, &annotation.VideoID,
		&annotation.Type, &annotation.Notes,
		&annotation.StartTime, &annotation.EndTime)

	if err != nil {
		return nil, err
	}
	return &annotation, nil
}

func (d *Database) ListAnnotations(videoID uuid.UUID) ([]*model.Annotation, error) {
	query := `SELECT * FROM annotations WHERE video_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := d.Client.QueryContext(ctx, query, videoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []*model.Annotation
	for rows.Next() {
		var result model.Annotation
		err = rows.Scan(&result.ID, &result.VideoID, &result.Type,
			&result.Notes, &result.StartTime, &result.EndTime)
		if err != nil {
			return nil, err
		}
		results = append(results, &result)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *Database) CreateAnnotation(newAnnotaton *model.Annotation) (*model.Annotation, error) {

	query := `INSERT INTO annotations (id, video_id, type, notes, start_time, end_time)
              VALUES ($1,$2,$3,$4,$5,$6) 
              RETURNING id`

	id := uuid.New()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := d.Client.QueryRowContext(ctx, query,
		id, newAnnotaton.VideoID,
		newAnnotaton.Type, newAnnotaton.Notes,
		newAnnotaton.StartTime, newAnnotaton.EndTime).Scan(&newAnnotaton.ID)

	if err != nil {
		return nil, err
	}
	return newAnnotaton, nil
}

func (d *Database) UpdateAnnotation(annotationID uuid.UUID, updatedAnno *model.AnnotationMetadata) error {

	query := `UPDATE annotations SET type = $2, notes = $3, start_time = $4, end_time = $5 WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := d.Client.ExecContext(ctx, query,
		annotationID, updatedAnno.Type,
		updatedAnno.Notes, updatedAnno.StartTime,
		updatedAnno.EndTime)

	if err != nil {
		return err
	}

	return nil
}

func (d *Database) DeleteAnnotation(annotationID uuid.UUID) error {
	query := `DELETE FROM annotations WHERE id = $1`

	result, err := d.Client.Exec(query, annotationID)

	if err != nil {
		return err
	}

	n, _ := result.RowsAffected()

	if n < 1 {
		return errors.New("not found")
	}

	return nil
}

func (d *Database) GetAnnotationVideoID(annotationID uuid.UUID) (*uuid.UUID, error) {
	query := `SELECT video_id FROM annotations WHERE annotations.id = $1`

	var videoID uuid.UUID

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := d.Client.QueryRowContext(ctx, query, annotationID).Scan(&videoID)
	if err != nil {
		return nil, err
	}

	return &videoID, err
}
