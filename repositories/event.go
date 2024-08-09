package repositories

import (
	"context"

	"github.com/satyam-jha-16/event-manager/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}
	res := r.db.Model(&models.Event{}).Order("updated_at desc").Find(&events)

	if res.Error != nil {
		return nil, res.Error
	}

	return events, res.Error
}

func (r *EventRepository) GetOne(ctx context.Context, id uint) (*models.Event, error) {
	event := &models.Event{}
	res := r.db.Model(event).Where("id = ?", id).First(event)

	if res.Error != nil {
		return nil, res.Error
	}

	return event, res.Error
}

func (r *EventRepository) Create(ctx context.Context, event *models.Event) (*models.Event, error) {
	res := r.db.Create(event)

	if res.Error != nil {
		return nil, res.Error
	}
	return event, nil
}

func (r *EventRepository) DeleteOne(ctx context.Context, id uint) error {
	res := r.db.Delete(&models.Event{}, id)
	return res.Error
}

func (r *EventRepository) UpdateOne(ctx context.Context, id uint, updateData map[string]interface{}) (*models.Event, error) {
	event := &models.Event{}
	upres := r.db.Model(event).Where("id = ?", id).Updates(updateData)

	if upres.Error != nil {
		return nil, upres.Error
	}

	getres := r.db.Model(event).Where("id = ?", id).First(event)
	if getres.Error != nil {
		return nil, getres.Error
	}
	return event, nil
}

func NewEventRepository(db *gorm.DB) models.EventRepository {
	return &EventRepository{db: db}
}
