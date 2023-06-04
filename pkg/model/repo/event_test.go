package repo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/errs"
	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/event"
	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/prototype"
)

func TestEventRepoInsert(t *testing.T) {
	ctx := context.Background()

	SetEventRepo(MockEventRepo{
		InsertFunc: func(ctx context.Context, e event.Event) errs.AppError {
			return nil
		},
	})
	defer SetEventRepo(nil)

	newEvent := prototype.PrototypeEvent()

	err := GetEventRepo().Insert(ctx, newEvent)
	assert.NoError(t, err)
}
