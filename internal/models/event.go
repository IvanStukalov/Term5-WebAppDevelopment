package models

import "time"

type Event struct {
	ID             int       `json:"event_id" gorm:"primaryKey;column:event_id;not null"`
	Name           string    `json:"name"`
	Status         string    `json:"status"`
	CreationDate   time.Time `json:"creation_date"`
	FormationDate  time.Time `json:"formation_date"`
	CompletionDate time.Time `json:"completion_date"`
	ModeratorID    int       `json:"moderator_id"`
	CreatorID      int       `json:"creator_id"`
}

type EventMsg struct {
	StarID    int `json:"star_id" gorm:"column:star_id"`
	CreatorID int `json:"creator_id" gorm:"column:creator_id;not null"`
}

type EventDetails struct {
	Event     Event  `json:"event"`
	StarsList []Star `json:"stars_list"`
}

type StarEvents struct {
	StarID  int `json:"star_id"`
	EventID int `json:"event_id"`
}
