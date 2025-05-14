package repository

import "context"

func (r *repository) count(ctx context.Context, sql string, args []any) (*int, error) {
	var n int
	if err := r.pgx.QueryRow(ctx, sql, args...).Scan(&n); err != nil {
		return nil, err
	}
	return &n, nil
}