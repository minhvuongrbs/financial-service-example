package campaign

import "time"

type Status int

const (
	StatusActive Status = iota
	StatusInactive
	StatusUnknown
)

type Campaign struct {
	Id int64

	Key  string
	Name string

	Status Status

	StartAt time.Time
	EndAt   time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCampaign(key, name string, status Status, startAt, endAt time.Time) (*Campaign, error) {
	return &Campaign{
		Id:      0,
		Key:     key,
		Name:    name,
		Status:  status,
		StartAt: startAt,
		EndAt:   endAt,
	}, nil
}
