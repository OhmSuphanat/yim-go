package infrastructure

import (
	"context"
	"log"

	"cloud.google.com/go/secretmanager/apiv1"
)

func NewSecretManagerClient(ctx context.Context) *secretmanager.Client {
	// Delete this line when you are ready to use the secret manager
	return nil

	secretCli, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("unable to create secret manager client: %v", err)
	}
	return secretCli
}