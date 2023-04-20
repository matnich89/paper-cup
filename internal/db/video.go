package db

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"papercup-test/internal/model"
)

func (d *Database) GetVideoLength(videoID uuid.UUID) (string, error) {
	query := `SELECT length FROM videos WHERE id = $1`

	var length string

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := d.Client.QueryRowContext(ctx,query, videoID).Scan(&length)

	if err != nil {
		return "", err
	}
	return length, nil
}

func (d *Database) CreateVideo(newVideo *model.Video) (*model.Video, error) {

	query := `INSERT INTO videos (id, name, url, length) VALUES ($1,$2,$3,$4)
              RETURNING id`

	id := uuid.New()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := d.Client.QueryRowContext(ctx, query, id, newVideo.Name, newVideo.URL, newVideo.Length).Scan(&newVideo.ID)

	if err != nil {
		return nil, err
	}
	return newVideo, nil
}

func (d *Database) DeleteVideo(videoID uuid.UUID) error {
	query := `DELETE FROM videos WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := d.Client.ExecContext(ctx,query, videoID)

	if err != nil {
		return err
	}

	n, _ := result.RowsAffected()

	if n < 1 {
		return errors.New("not found")
	}

	return nil
}
