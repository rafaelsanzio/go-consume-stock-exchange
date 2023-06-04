package repo

import (
	"context"

	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/errs"
	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/event"
)

type MockEventRepo struct {
	event.EventRepo
	InsertFunc func(ctx context.Context, e event.Event) errs.AppError
}

func (m MockEventRepo) Insert(ctx context.Context, e event.Event) errs.AppError {
	if m.InsertFunc != nil {
		return m.InsertFunc(ctx, e)
	}
	return m.EventRepo.Insert(ctx, e)
}
