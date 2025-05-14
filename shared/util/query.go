package util

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type queryKey struct{}

type TxnObjKey struct{}

type QueryUtil struct {
	pool *pgxpool.Pool
	scanApi *pgxscan.API
}

type Args interface {
	pgx.NamedArgs | []any
}

func NewQueryUtil(pool *pgxpool.Pool, scanAPi *pgxscan.API) *QueryUtil {
	return &QueryUtil{
		pool: pool,
		scanApi: scanAPi,
	}
}

func (q *QueryUtil) SetPoolTx(ctx context.Context) context.Context {
	if tx, ok := ctx.Value(TxnObjKey{}).(*pgx.Tx); ok {
		return context.WithValue(ctx, queryKey{}, tx)
	}
	return context.WithValue(ctx, queryKey{}, q.pool)
}

func (q *QueryUtil) Exec(ctx context.Context, querystring string, args []any) (err error) {
	switch poolTx := ctx.Value(queryKey{}).(type) {
	case *pgxpool.Pool:
		_, err = poolTx.Exec(ctx, querystring, args...)
	case pgx.Tx:
		_, err = poolTx.Exec(ctx, querystring, args...)
	}
	return err
}

func (q *QueryUtil) ExecNameArgs(ctx context.Context, querystring string, args pgx.NamedArgs) (err error) {
	switch poolTx := ctx.Value(queryKey{}).(type) {
	case *pgxpool.Pool:
		_, err = poolTx.Exec(ctx, querystring, args)
	case pgx.Tx:
		_, err = poolTx.Exec(ctx, querystring, args)
	}
	return err
}

func (q *QueryUtil) ExecWithRowAffected(ctx context.Context, querystring string, args any) (int, error) {
	var (
		commandTag 	pgconn.CommandTag
		err	   		error
		isNamedArgs	bool
		namedArgs 	pgx.NamedArgs
		arrArgs		[]any
	)

	switch parsedArgs := args.(type) {
		case pgx.NamedArgs:
			isNamedArgs = true
			namedArgs = parsedArgs
		case []any:
			arrArgs = parsedArgs
		default:
			return 0, errors.New("invalid argument type")
	}

	switch poolTx := ctx.Value(queryKey{}).(type) {
		case *pgxpool.Pool:
			if isNamedArgs {
				commandTag, err = poolTx.Exec(ctx, querystring, namedArgs)
				return int(commandTag.RowsAffected()), err
			}
			commandTag, err = poolTx.Exec(ctx, querystring, arrArgs...)
		case pgx.Tx:
			if isNamedArgs {
				commandTag, err = poolTx.Exec(ctx, querystring, namedArgs)
				return int(commandTag.RowsAffected()), err
			}
		default:
			return 0, errors.New("unsupported context type")
	}
	return int(commandTag.RowsAffected()), err
}

func (q *QueryUtil) QueryRow(ctx context.Context, dest interface{}, querystring string, args []any) error {
	var (
		row pgx.Rows
		err error
	)

	switch poolTx := ctx.Value(queryKey{}).(type) {
	case *pgxpool.Pool:
		row, err = poolTx.Query(ctx, querystring, args...)
	case pgx.Tx:
		row, err = poolTx.Query(ctx, querystring, args...)
	default:
		return errors.New("unsupported context type")
	}

	if err != nil {
		return err
	}
	defer row.Close()

	if err := q.scanApi.ScanOne(dest, row); err != nil {
		return err
	}
	return nil
}

func (q *QueryUtil) Query(ctx context.Context, dest interface{}, querystring string, args []any) error {
	var (
		rows pgx.Rows
		err  error
	)

	switch poolTx := ctx.Value(queryKey{}).(type) {
	case *pgxpool.Pool:
		rows, err = poolTx.Query(ctx, querystring, args...)
	case pgx.Tx:
		rows, err = poolTx.Query(ctx, querystring, args...)
	default:
		return errors.New("unsupported context type")
	}

	if err != nil {
		return err
	}
	defer rows.Close()

	if err := q.scanApi.ScanAll(dest, rows); err != nil {
		return err
	}
	return nil
}

func (q *QueryUtil) BulkUpsertArgsBuild(args pgx.NamedArgs, lineStm []string, masterColumn []string, index int, mapper map[string]any) {
	placeHolder := strconv.Itoa(index + 1)
	lineHolders := make([]string, len(masterColumn))
	for i, column := range masterColumn {
		argsKey := column + "_" + placeHolder
		args[argsKey] = mapper[column]
		lineHolders[i] = fmt.Sprint("@%s", argsKey)
	}

	lineStm[index] = fmt.Sprintf("(%s)", strings.Join(lineHolders, ","))
}