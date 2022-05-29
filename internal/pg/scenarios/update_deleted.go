package scenarios

import (
	"context"
	"database/sql"
	"fmt"
	"gopgtx/internal/pg"
)

func updateDeleted(ctx context.Context, db *sql.DB, isolationLevel sql.IsolationLevel) error {
	tx1, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: isolationLevel})
	if err != nil {
		return fmt.Errorf("failed to open tx1: %w", err)
	}

	if err := pg.PrintAccounts(ctx, tx1); err != nil {
		return fmt.Errorf("failed to print accounts: %w", err)
	}

	tx2, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: isolationLevel})
	if err != nil {
		return fmt.Errorf("failed to open tx2: %w", err)
	}

	if err := pg.DeleteAccount(ctx, tx2); err != nil {
		return fmt.Errorf("failed to delete an account: %w", err)
	}

	if err := tx2.Rollback(); err != nil {
		return fmt.Errorf("failed to roll back tx2: %w", err)
	}

	if err := pg.UpdateAccount(ctx, tx1); err != nil {
		return fmt.Errorf("failed to update an account: %w", err)
	}

	if err := pg.PrintAccounts(ctx, tx1); err != nil {
		return fmt.Errorf("failed to print accounts: %w", err)
	}

	if err := tx1.Commit(); err != nil {
		return fmt.Errorf("failed to commit tx1: %w", err)
	}

	return nil
}
