package repositories

import (
	"context"
	"time"

	"github.com/satyam-jha-16/event-manager/models"
)

type EventRepository struct {
	db any
}


func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error){
	events := []*models.Event{}
	events = append(events, &models.Event{
		ID: "asf123",
		Name: "star wars",
		Location: "galaxy far far away",
		Date: time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	
	return events, nil
}

func (r *EventRepository) GetOne(ctx context.Context, id string) (*models.Event, error){
	return nil , nil
}

func (r *EventRepository) Create(ctx context.Context, event *models.Event) (*models.Event, error){
	return nil , nil
}

func NewEventRepository(db any) models.EventRepository{
	return &EventRepository{db : db}
}