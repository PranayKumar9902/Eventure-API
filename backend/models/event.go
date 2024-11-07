package models

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID                    uint      `json:"id" gorm:"primaryKey"`
	Name                  string    `json:"name"`
	Location              string    `json:"location"`
	TotalTicketsPurchased int64     `json:"totalTicketsPurchased"`
	TotalTicketsEntered   int64     `json:"totalTicketsEntered"`
	Date                  time.Time `json:"date"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
}

type EventRepository interface {
	GetMany(ctx context.Context) ([]*Event, error)
	GetOne(ctx context.Context, eventId uint) (*Event, error)
	CreateOne(ctx context.Context, event Event) (*Event, error)
	UpdatedOne(ctx context.Context, eventId uint, updateData map[string]interface{}) (*Event, error)
	DeleteOne(ctx context.Context, eventId uint) error
}

func (event *Event) AfterFind(db *gorm.DB) (err error) {

	baseQuery := db.Model(&Ticket{}).Where("event_id = ?", event.ID)

	if err := baseQuery.Count(&event.TotalTicketsPurchased).Error; err != nil {
		return err
	}

	if err := baseQuery.Where("entered = ?", true).Count(&event.TotalTicketsEntered).Error; err != nil {
		return err
	}

	return nil
}
