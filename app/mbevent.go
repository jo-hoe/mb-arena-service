package app

import (
	"time"
)

type MBEvent struct {
	Name       string
	Link       string
	PictureUrl string
	Start      time.Time
}
