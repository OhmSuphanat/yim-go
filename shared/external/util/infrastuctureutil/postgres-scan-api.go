package infrastuctureutil

import (
	"context"
	"fmt"
	"log"
	"time"
	baseproperty "yim-go/shared/external/configuration/base-property"

	"github.com/exaring/otelpgx"
	"github.com/georgysavva/scany/v2/dbscan"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)


type PostgresConfig struct {
	// ConnUri: "host=localhost port=5432 database=postgres user=postgres password=xxx"
	ConnUri string `envconfig:"POSTGRES_CONN_URI" default:"host=localhost port=5432 database=postgres user=postgres password=%s"`
	ConnString string

	// MaxConnLifetime is the duration since creation after which a connetion will be automatically closed.
	MaxConnLifetime baseproperty.Duration `envconfig:"POSTGRES_MAX_CONN_LIFETIME" default:"1h"`

	// MaxConnIdleTime is the furation after which an idle connection will be automatically closed by the health check.
	MaxConnIdleTime baseproperty.Duration `envconfig:"POSTGRES_MAX_CONN_IDLE_TIME" default:"30m"`

	// MaxConns is the maximum size of the pool. The default is the greater of 4 or runtime.NumCPU();
	MaxConns int32 `envconfig:"POSTGRES_MAX_CONNS" default:"4"`

	// MinConns iss the minimum size of the pool. After connection closes, the pool might dip below MinConns.
	// A low number of MinConns might mean the pool is empty after MaxConnLifetime until the health check has a chance to create new connections.
	MinConns int32 `envconfig:"POSTGRES_MIN_CONNS" default:"0"`
}

func (p *PostgresConfig) SetConnPassword(password string) {
	p.ConnUri = fmt.Sprintf(p.ConnUri, password)
}

func NewScanApi() *pgxscan.API {
	scnaApi, err := pgxscan.NewDBScanAPI(dbscan.WithAllowUnknownColumns(true))
	if err != nil {
		log.Fatalf("error create dbscanapi: %v", err)
	}

	api, err := pgxscan.NewAPI(scnaApi)
	if err != nil {
		log.Fatalf("error create sqlscanner api: %v", err)
	}

	return api
}

func NewPostgresWithScanApi(ctx context.Context, pgCfg PostgresConfig) (*pgxpool.Pool, *pgxscan.API, error) {
	cfg, err := pgxpool.ParseConfig(pgCfg.ConnString)
	if err != nil {
		log.Fatalf("unable to parse postgres connection uri: %v", err)
	}

	cfg.ConnConfig.Tracer = otelpgx.NewTracer()
	if int64(pgCfg.MaxConnLifetime) > 0 {
		cfg.MaxConnLifetime = time.Duration(pgCfg.MaxConnLifetime)
	}
	if pgCfg.MaxConnIdleTime > 0 {
		cfg.MaxConnIdleTime = time.Duration(pgCfg.MaxConnIdleTime)
	}
	if pgCfg.MaxConns > 0 {
		cfg.MaxConns = pgCfg.MaxConns
	}
	if pgCfg.MinConns > 0 {
		cfg.MinConns = pgCfg.MinConns
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to create pg connection: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, nil, fmt.Errorf("unable to ping db: %v", err)
	}

	scanApi := NewScanApi()
	return pool, scanApi, nil
}