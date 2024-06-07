package app

import (
	"time"
)

type Event struct {
	Name       string    `json:"name"`
	Link       string    `json:"link"`
	PictureUrl string    `json:"pictureUrl"`
	Start      time.Time `json:"start"`
	Location   string    `json:"location"`
}
