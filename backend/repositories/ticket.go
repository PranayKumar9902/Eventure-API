package repositories

import (
	"context"

	"github.com/pranay/ticket-booking-app/models"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func (r *TicketRepository) GetMany(ctx context.Context, userId uint) ([]*models.Ticket, error) {

	tickets := []*models.Ticket{}
	ticket := &models.Ticket{}

	if err := r.db.Model(ticket).Where("user_id = ?", userId).Preload("Event").Order("updated_at desc").Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *TicketRepository) GetOne(ctx context.Context, userId uint, ticketId uint) (*models.Ticket, error) {

	ticket := &models.Ticket{}

	res := r.db.Model(ticket).Where("id = ?", ticketId).Where("user_id = ?", userId).Preload("Event").First(ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil
}

func (r *TicketRepository) CreateOne(ctx context.Context, userId uint, ticket models.Ticket) (*models.Ticket, error) {

	ticket.UserID = userId
	if err := r.db.Create(&ticket).Error; err != nil {
		return nil, err
	}
	return r.GetOne(ctx, userId, ticket.ID)
}

func (r *TicketRepository) UpdatedOne(ctx context.Context, userId uint, ticketId uint, updateData map[string]interface{}) (*models.Ticket, error) {

	ticket := &models.Ticket{}

	updateRes := r.db.Model(ticket).Where("id = ?", ticketId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	return r.GetOne(ctx, userId, ticketId)
}

func NewTicketRepository(db *gorm.DB) models.TicketRepository {
	return &TicketRepository{
		db: db,
	}
}
