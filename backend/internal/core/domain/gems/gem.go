package gems

import (
	"time"
)

type Gem struct {
	ID         int64
	Message    string
	Latitude   float64
	Longitude  float64
	CreatorId  int64
	CreatedAt  time.Time
	ReceiverId int64
	ReceivedAt *time.Time
}
