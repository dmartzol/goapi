package goapi

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID          uuid.UUID  `json:"Id"`
	CreatedTime time.Time  `db:"create_time"`
	UpdatedTime *time.Time `db:"update_time"`
}

func NewModel() *Model {
	now := time.Now().UTC()
	return &Model{
		ID:          uuid.New(),
		CreatedTime: now,
		UpdatedTime: &now,
	}
}
