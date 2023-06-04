package prototype

import "github.com/rafaelsanzio/go-consume-stock-exchange/pkg/event"

func PrototypeEvent() event.Event {
	return event.Event{
		ID:            "1",
		Ticker:        "VWX",
		Price:         150.5,
		Change:        2,
		PercentChange: 1.25,
	}
}
