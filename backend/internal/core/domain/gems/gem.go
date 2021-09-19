package gems

import (
	"time"
)

type GemColor int64

const (
	Purple GemColor = iota
	Pink
	Blue
	Black
	Yellow
	Green
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
	Color      GemColor
}
