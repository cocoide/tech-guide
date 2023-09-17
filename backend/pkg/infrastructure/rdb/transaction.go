package rdb

import (
	"context"
	"github.com/cocoide/tech-guide/key"
)

func (r *Repository) DoInTx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx := r.db.Begin()
	c := context.WithValue(ctx, key.Transaction, tx)
	var done bool
	defer func() {
		if !done {
			_ = tx.Rollback()
		}
	}()
	if err := fn(c); err != nil {
		return err
	}
	if ctx.Err() == context.DeadlineExceeded {
		return context.DeadlineExceeded
	}
	done = true
	return tx.Commit().Error
}
