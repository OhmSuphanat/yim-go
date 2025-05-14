package infrastructure

import (
	"context"
	"yim-go/shared/external/util/infrastuctureutil"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(ctx context.Context, postgresConfig infrastuctureutil.PostgresConfig) (*pgxpool.Pool, *pgxscan.API) {
	pgx, scanapi, err := infrastuctureutil.NewPostgresWithScanApi(ctx, postgresConfig)
	if err != nil {
		panic(err)
	}
	return pgx, scanapi
}