package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/applog"
	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/config"
	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/config/key"
	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/consumer"
	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/errs"
	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/store"
)

func main() {
	key.LoadEnvVars()

	store.GetStore()

	mongoURI, err_ := config.Value(key.MongoURI)
	if err_ != nil {
		_ = err_.Annotatef(applog.Log, "unable to get mongo config: %v", err_)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		_ = errs.ErrMongoConnect.Throwf(applog.Log, errs.ErrFmt, err)
	}

	ctx, cancFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancFunc()

	err = client.Connect(ctx)
	if err != nil {
		_ = errs.ErrMongoConnect.Throwf(applog.Log, errs.ErrFmt, err)
	}

	defer func() {
		err = client.Disconnect(ctx)
		if err != nil {
			_ = errs.ErrMongoConnect.Throwf(applog.Log, errs.ErrFmt, err)
		}
	}()

	err = consumer.Consumer("stockexchange")
	if err != nil {
		log.Fatal(err)
	}
}
