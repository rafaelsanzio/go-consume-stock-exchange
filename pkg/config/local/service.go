package local

import (
	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/config/key"
	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/errs"
)

type Service struct{}

func (s Service) Value(key key.Key) (string, errs.AppError) {
	return Values[key], nil
}
