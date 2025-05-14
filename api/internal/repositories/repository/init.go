package repository

import (
	"yim-go/api/internal/core/port"
	"yim-go/shared/util"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	pgx	*pgxpool.Pool
	queryUtil *util.QueryUtil
	sqlbuilder sqlbuilder.Flavor
}

func New(pgx *pgxpool.Pool, scanApi *pgxscan.API, sqlbuilder sqlbuilder.Flavor) port.Repository{
	return &repository{
		pgx: pgx,
		queryUtil: util.NewQueryUtil(pgx, scanApi),
		sqlbuilder: sqlbuilder,
	}
}