package local

import (
	"os"

	"github.com/rafaelsanzio/go-consume-stock-exchange/pkg/config/key"
)

var Values = map[key.Key]string{
	key.MongoDBName:     getDefaultOrEnvVar("stockexchange", "MONGO_DATABASE"),
	key.MongoDBPassword: getDefaultOrEnvVar("root", "MONGO_PASWORD"),
	key.MongoDBUsername: getDefaultOrEnvVar("root", "MONGO_USERNAME"),
	key.MongoURI:        getDefaultOrEnvVar("mongodb://root:root@mongo:27017/?connect=direct", "MONGO_URI"),

	key.NatsURL:  getDefaultOrEnvVar("nats://go-consume-stock-exchange_nats_1", "NATS_URL"),
	key.NatsPort: getDefaultOrEnvVar("4222", "NATS_PORT"),

	key.AppPort: getDefaultOrEnvVar("9092", "APP_PORT"),
}

// Some of the db fields are set via env var in the makefile, so this optionally uses those to prevent test failures in jenkins
func getDefaultOrEnvVar(dfault, envVar string) string {
	val := os.Getenv(envVar)
	if val != "" {
		return val
	}
	return dfault
}
