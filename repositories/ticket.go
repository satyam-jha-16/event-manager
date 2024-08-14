package repositories

import (
	"context"
	"fmt"

	"github.com/satyam-jha-16/event-manager/models"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func (r *TicketRepository) GetMany(ctx context.Context, userId uint) ([]*models.Ticket, error) {
    tickets := []*models.Ticket{}

    res := r.db.Model(&models.Ticket{}).Where("user_id = ?", userId).Preload("Event").Order("updated_at desc").Find(&tickets)

    if res.Error != nil {
        return nil, res.Error
    }

    return tickets, nil
}




func (r *TicketRepository) GetOne(ctx context.Context, userId uint, id uint) (*models.Ticket, error) {
    ticket := &models.Ticket{}

    res := r.db.Model(&models.Ticket{}).Where("id = ?", id).Where("user_id = ?", userId).Preload("Event").First(ticket)

    if res.Error != nil {
        return nil, res.Error
    }
    return ticket, nil
}

func (r *TicketRepository) Create(ctx context.Context, userId uint, ticket *models.Ticket) (*models.Ticket, error) {
	ticket.UserID = userId
    if err := r.db.First(&models.Event{}, ticket.EventID).Error; err != nil {
        return nil, fmt.Errorf("invalid event ID: %v", err)
    }

    res := r.db.Model(&models.Ticket{}).Create(ticket)
    if res.Error != nil {
        return nil, res.Error
    }

    return r.GetOne(ctx, userId, ticket.ID)

}


func (r *TicketRepository) DeleteOne(ctx context.Context, userId uint, id uint) error {

	res := r.db.Delete(&models.Ticket{}, id)
	return res.Error
}

func (r *TicketRepository) UpdateOne(ctx context.Context,userId uint, id uint, updateData map[string]interface{}) (*models.Ticket, error) {
	ticket := &models.Ticket{}
	upres := r.db.Model(ticket).Where("id = ?", id).Updates(updateData)

	if upres.Error != nil {
		return nil, upres.Error
	}

	getres := r.db.Preload("Event").Model(ticket).Where("id = ?", id).First(ticket)
	if getres.Error != nil {
		return nil, getres.Error
	}

	return ticket, nil
}

func NewTicketRepository(db *gorm.DB) models.TicketRepository {
	return &TicketRepository{db: db}
}
