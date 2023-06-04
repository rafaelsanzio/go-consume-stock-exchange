package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"

	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/applog"
	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/config"
	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/config/key"
	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/errs"
	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/event"
	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/model/repo"
)

func Consumer(topic string) error {
	ctx := context.Background()

	natsURL, err_ := config.Value(key.NatsURL)
	if err_ != nil {
		_ = errs.ErrGettingEnvNatsURL.Throwf(applog.Log, errs.ErrFmt, err_)
	}

	natsPort, err_ := config.Value(key.NatsPort)
	if err_ != nil {
		_ = errs.ErrGettingEnvNatsPort.Throwf(applog.Log, errs.ErrFmt, err_)
	}

	nc, err := nats.Connect(fmt.Sprintf("%s:%s", natsURL, natsPort))
	if err != nil {
		_ = errs.ErrNatsConnection.Throwf(applog.Log, errs.ErrFmt, err)
		return nil
	}

	defer nc.Close()

	// Build an event channel to the topic
	ch := make(chan *nats.Msg, 64)
	sub, err := nc.ChanSubscribe(topic, ch)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	// Wait for new messages and insert into DB
	for msg := range ch {
		log.Printf("Read message from NATS: %s\n", string(msg.Data))

		err := insertMessage(ctx, msg.Data)
		if err != nil {
			return err
		}
	}

	return nil
}

func insertMessage(ctx context.Context, message []byte) errs.AppError {
	event := &event.Event{}
	err := json.Unmarshal(message, event)
	if err != nil {
		return errs.ErrUnmarshalingJson.Throwf(applog.Log, errs.ErrFmt, err.Error())
	}

	err = repo.GetEventRepo().Insert(ctx, *event)
	if err != nil {
		return errs.ErrMongoInsertOne.Throwf(applog.Log, errs.ErrFmt, err.Error())
	}

	return nil
}
