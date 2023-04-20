package model

import (
	"github.com/google/uuid"
)

type Video struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	URL    string    `json:"url"`
	Length string    `json:"length"`
}

type Annotation struct {
	ID      uuid.UUID `json:"id"`
	VideoID uuid.UUID `json:"video_id"`
	AnnotationMetadata
}

type AnnotationMetadata struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Type      string `json:"type"`
	Notes     string `json:"notes"`
}

type Credentials struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}
