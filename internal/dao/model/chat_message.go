package model

import "time"

type ChatMessage struct {
	Id string
	Content string
	CreatedAt time.Time
	CreatedBy string
}
