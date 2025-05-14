package property

import (
	"context"
	configs "yim-go/shared/external/configuration/backoffice-property"
	baseproperty "yim-go/shared/external/configuration/base-property"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
)

func GetBaseProperty() baseproperty.ServerProperties {
	return baseproperty.ServerProperties(configs.Get())
}

func setPostgresSecret(ctx context.Context, secretCli *secretmanager.Client) {
	if len(cfg.Secret.PostgresPasswordSecret) == 0 {
		password := "postgres"
		cfg.Postgres.PostgresConfig.SetConnPassword(password)
		cfg.Postgres.PostgresConfig.ConnString = cfg.Postgres.PostgresConfig.ConnUri
	}
}
