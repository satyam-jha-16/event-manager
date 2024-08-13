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

func (r *TicketRepository) GetMany(ctx context.Context) ([]*models.Ticket, error) {
    tickets := []*models.Ticket{}
    
    res := r.db.Preload("Event").Order("updated_at desc").Find(&tickets)
    
    if res.Error != nil {
        return nil, res.Error
    }
    
    return tickets, nil 
}

func (r *TicketRepository) GetOne(ctx context.Context, id uint) (*models.Ticket, error) {
    ticket := &models.Ticket{}
    
    res := r.db.Preload("Event").Where("id = ?", id).First(ticket)
    
    if res.Error != nil {
        return nil, res.Error
    }
    return ticket, nil
}

func (r *TicketRepository) Create(ctx context.Context, ticket *models.Ticket) (*models.Ticket, error) {
    // Verify the associated Event exists
    if err := r.db.First(&models.Event{}, ticket.EventID).Error; err != nil {
        return nil, fmt.Errorf("invalid event ID: %v", err)
    }
    
    // Create the Ticket
    res := r.db.Create(ticket)
    if res.Error != nil {
        return nil, res.Error
    }
    
    // Reload the ticket with the associated event
    if err := r.db.Preload("Event").First(ticket, ticket.ID).Error; err != nil {
        return nil, err
    }
    
    return ticket, nil
}


func (r *TicketRepository) DeleteOne(ctx context.Context, id uint) error {

	res := r.db.Delete(&models.Ticket{}, id)
	return res.Error
}

func (r *TicketRepository) UpdateOne(ctx context.Context, id uint, updateData map[string]interface{}) (*models.Ticket, error) {
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
