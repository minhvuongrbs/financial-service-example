package campaign

import (
	"fmt"
	"time"
)

type Status int

const (
	StatusActive Status = iota
	StatusInactive
	StatusUnknown
)

type Campaign struct {
	Id int64

	Name   string
	Status Status

	Metadata Metadata

	CreatedAt time.Time
	UpdatedAt time.Time
}

const (
	minNameLength = 5
)

func NewCampaign(id int64, name string, status Status, md Metadata) (*Campaign, error) {
	if len(name) < minNameLength {
		return nil, fmt.Errorf("invalid name length")
	}
	if status == StatusUnknown {
		return nil, fmt.Errorf("invalid status: %v", status)
	}

	return &Campaign{
		Id:       id,
		Name:     name,
		Status:   status,
		Metadata: md,
	}, nil
}
