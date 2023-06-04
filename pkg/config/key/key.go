package key

import (
	"github.com/joho/godotenv"

	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/applog"
	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/errs"
)

type Key struct {
	Name   string
	Secure bool
	Provider
}

type Provider string

var (
	ProviderStore  = Provider("store")
	ProviderEnvVar = Provider("env")
)

var (
	MongoDBName     = Key{Name: "MONGO_DATABASE", Secure: false, Provider: ProviderStore}
	MongoDBPassword = Key{Name: "MONGO_PASSWORD", Secure: true, Provider: ProviderStore}
	MongoDBUsername = Key{Name: "MONGO_USERNAME", Secure: false, Provider: ProviderStore}
	MongoURI        = Key{Name: "MONGO_URI", Secure: false, Provider: ProviderStore}

	NatsURL  = Key{Name: "NATS_URL", Secure: false, Provider: ProviderEnvVar}
	NatsPort = Key{Name: "NATS_PORT", Secure: false, Provider: ProviderEnvVar}
	AppPort  = Key{Name: "APP_PORT", Secure: false, Provider: ProviderEnvVar}
)

func LoadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		_ = errs.ErrGettingEnv.Throwf(applog.Log, errs.ErrFmt, err)
	}
}
