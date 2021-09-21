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
	ID         int64      `db:"id" goqu:"skipinsert"`
	Message    string     `db:"message"`
	Latitude   float64    `db:"latitude"`
	Longitude  float64    `db:"longitude"`
	CreatorId  int64      `db:"creator_id"`
	CreatedAt  time.Time  `db:"created_at" goqu:"skipinsert"`
	ReceiverId int64      `db:"receiver_id"`
	ReceivedAt *time.Time `db:"received_at"`
	Color      GemColor   `db:"color"`
	Attachment []byte     `db:"attachment"`
}
