package repositories

import (
	"context"

	"github.com/pranay/ticket-booking-app/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {

	events := []*models.Event{}

	if err := r.db.Find(&events).Order("updated_at desc").Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) GetOne(ctx context.Context, eventId uint) (*models.Event, error) {

	event := &models.Event{}

	if err := r.db.Find(&event, "id = ?", eventId).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event models.Event) (*models.Event, error) {

	if err := r.db.Create(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *EventRepository) UpdatedOne(ctx context.Context, eventId uint, updateData map[string]interface{}) (*models.Event, error) {

	event := &models.Event{}

	if err := r.db.Model(&event).Where("id = ?", eventId).Updates(updateData).Error; err != nil {
		return nil, err
	}

	if err := r.db.Find(&event, "id = ?", eventId).Error; err != nil {
		return nil, err
	}

	return event, nil
}

func (r *EventRepository) DeleteOne(ctx context.Context, eventId uint) error {

	event := &models.Event{}

	if err := r.db.Find(&event, "id = ?", eventId).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&event).Error; err != nil {
		return err
	}

	return nil
}

func NewEventRepository(db *gorm.DB) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}
