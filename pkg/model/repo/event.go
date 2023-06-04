package repo

import (
	"context"
	"time"

	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/errs"
	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/event"
	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/store"
)

const (
	EventCollection = "events"
)

type eventRepo struct {
	store store.Store
}

var eventRepoSingleton event.EventRepo

func GetEventRepo() event.EventRepo {
	if eventRepoSingleton == nil {
		return getEventRepo()
	}
	return eventRepoSingleton
}

func getEventRepo() *eventRepo {
	s := store.GetStore()
	return &eventRepo{s}
}

func SetEventRepo(repo event.EventRepo) {
	eventRepoSingleton = repo
}

func (repo eventRepo) Insert(ctx context.Context, e event.Event) errs.AppError {
	e.Created = time.Now()
	_, err := repo.store.InsertOne(ctx, EventCollection, &e)
	return err
}
