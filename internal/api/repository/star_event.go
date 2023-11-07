package repository

import (
	"errors"
	"github.com/IvanStukalov/Term5-WebAppDevelopment/internal/models"
)

func (r *Repository) RemoveStarFromEvent(creatorId int, starId int) (models.EventDetails, error) {
	var event models.Event
	r.db.Where("status = ?", models.StatusCreated).Where("creator_id = ?", creatorId).First(&event)

	if event.ID == 0 {
		return models.EventDetails{}, errors.New("no such request")
	}

	var starEvent models.StarEvents
	err := r.db.Where("event_id = ? AND star_id = ?", event.ID, starId).First(&starEvent).Error
	if err != nil {
		return models.EventDetails{}, errors.New("такой звезды нет в данном событии")
	}

	err = r.db.Where("event_id = ? AND star_id = ?", event.ID, starId).
		Delete(models.StarEvents{}).Error

	if err != nil {
		return models.EventDetails{}, err
	}

	return r.GetEventByID(event.ID)
}
