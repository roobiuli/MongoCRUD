package comment

import "time"

type Response struct {
	UUID      string    `json:"uuid"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}