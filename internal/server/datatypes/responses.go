package datatypes

import "time"

type ChatMessageResponse struct {
	Id string `json:"id"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
}
