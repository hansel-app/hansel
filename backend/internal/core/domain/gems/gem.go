package gems

import (
	"time"
)

type GemColor int64

const (
	Purple GemColor = 0
	Pink   GemColor = 1
	Blue   GemColor = 2
	Black  GemColor = 3
	Yellow GemColor = 4
	Green  GemColor = 5
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
	Photo      []byte
}
