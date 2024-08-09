package models

import (
	"context"
	"time"
)

type Event struct {
	ID uint `json:"id" gorm:"primarykey"`
	Name string `json:"name"`
	Location string `json:"location"`
	Date time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


type EventRepository interface {
	GetMany(ctx context.Context) ([]*Event, error)
	GetOne(ctx context.Context, eventId uint) (*Event, error)
	Create(ctx context.Context, event *Event) (*Event, error)
	UpdateOne(ctx context.Context, eventId uint, updateData map[string]interface{} ) (*Event, error)
	DeleteOne(ctx context.Context, eventId uint) error
}
