package event

import (
	"context"
	"time"

	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/errs"
)

type EventRepo interface {
	Insert(ctx context.Context, e Event) errs.AppError
}

type Event struct {
	ID            string `bson:"_id"`
	Ticker        string
	Price         float64
	Change        float64
	PercentChange float64
	Created       time.Time
}

func (e Event) GetID() string {
	return e.ID
}

func (e *Event) SetID(id string) {
	e.ID = id
}
