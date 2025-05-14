package property

import (
	"context"
	"log"
	configs "yim-go/shared/external/configuration/backoffice-property"
	"yim-go/shared/external/util/infrastuctureutil"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/kelseyhightower/envconfig"
)

type serverProperties struct {
	configs.ServerProperties
	ServerReflection bool `envconfig:"SERVER_REFLECTION"`
}


type postgres struct {
	PostgresConfig infrastuctureutil.PostgresConfig
}

type secret struct {
	PostgresPasswordSecret string `envconfig:"POSTGRES_PASSWORD_SECRET" default:""`
}

type config struct {
	Server serverProperties
	Postgres postgres
	Secret secret
}

var cfg config

func Init(ctx context.Context, secretCli *secretmanager.Client) {
	configs.InitServerProperties()
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("read env error: %s", err.Error())
	}

	// set postgres connection string
	setPostgresSecret(ctx, secretCli)
}

func Get() config {
	return cfg
}