package app

import (
	"time"
)

type MBEvent struct {
	Name       string    `json:"name"`
	Link       string    `json:"link"`
	PictureUrl string    `json:"pictureUrl"`
	Start      time.Time `json:"start"`
}
