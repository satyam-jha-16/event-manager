package models

import (
	"context"
	"time"
)

type Ticket struct {
    ID        uint      `json:"id" gorm:"primarykey"`
    EventID   uint      `json:"event_id"`
    UserID    uint      `json:"user_id" gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    Event     Event     `json:"event" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    Entered   bool      `json:"entered" default:"false"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}


type TicketRepository interface {
	GetMany(ctx context.Context, userId uint) ([]*Ticket, error)
	GetOne(ctx context.Context,userId uint, ticketID uint) (*Ticket, error)
	Create(ctx context.Context,userId uint, ticket *Ticket) (*Ticket, error)
	UpdateOne(ctx context.Context,userId uint, ticketID uint, updateData map[string]interface{}) (*Ticket, error)
}

type ValidateTicket struct{
	TicketID uint `json:"ticket_id"`
	OwnerID uint `json:"owner_id"`
}