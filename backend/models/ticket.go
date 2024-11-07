package models

import (
	"context"
	"time"
)

type Ticket struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	EventID   uint      `json:"eventId"`
	UserID    uint      `json:"userId" gorm:"foreignKey:UserID;OnDelete:CASCADE,OnUpdate:CASCADE"`
	Event     Event     `json:"event" gorm:"foreignKey:EventID;OnDelete:CASCADE,OnUpdate:CASCADE"`
	Entered   bool      `json:"entered" default:"false"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TicketRepository interface {
	GetMany(ctx context.Context, userId uint) ([]*Ticket, error)
	GetOne(ctx context.Context, userId uint, ticketId uint) (*Ticket, error)
	CreateOne(ctx context.Context, userId uint, ticket Ticket) (*Ticket, error)
	UpdatedOne(ctx context.Context, userId uint, ticketId uint, updateData map[string]interface{}) (*Ticket, error)
}

type ValidateTicket struct {
	TicketID uint `json:"ticketId"`
	OwnerID  uint `json:"ownerId"`
}
