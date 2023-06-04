package config

import (
	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/config/key"
	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/errs"
)

type Service interface {
	Value(key.Key) (string, errs.AppError)
}
