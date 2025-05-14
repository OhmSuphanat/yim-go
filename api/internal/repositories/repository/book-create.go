package repository

import (
	"context"
	"fmt"
	"yim-go/api/internal/core/domain"
	"yim-go/shared/constant"

	"github.com/jackc/pgx/v5"
)

func (r *repository) CreateBook(ctx context.Context, in domain.CreateBookRequest) error {

	ctx = r.queryUtil.SetPoolTx(ctx)

	query := fmt.Sprintf(`
		INSERT INTO %s (
			title,
			author,
			description,
			created_at,
			updated_at
		) VALUES (
		 @title,
		 @author,
		 @description,
		 @created_at,
		 @updated_at
		)`, constant.BookTable)
	
		args := pgx.NamedArgs{
			"title":       in.Title,
			"author":      in.Author,
			"description": in.Description,
			"created_at":  in.CreatedAt,
			"updated_at":  in.UpdatedAt,
		}

		if err := r.queryUtil.ExecNameArgs(ctx, query, args); err != nil {
			return err
		}
		return nil
}