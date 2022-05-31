package scenarios

import (
	"context"
	"database/sql"
	"fmt"

	"gopgtx/internal/pg"
)

func selectCommitted(db *sql.DB, isolationLevel sql.IsolationLevel) error {
	tx1, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: isolationLevel})
	if err != nil {
		return fmt.Errorf("failed to open tx1: %w", err)
	}

	if err := pg.PrintAccounts(tx1, `SELECT * FROM accounts WHERE balance > 50;`); err != nil {
		return fmt.Errorf("failed to print accounts: %w", err)
	}

	tx2, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: isolationLevel})
	if err != nil {
		return fmt.Errorf("failed to open tx2: %w", err)
	}

	if err := pg.UpdateAccount(tx2, `UPDATE accounts SET balance = 51.0 WHERE balance < 50.0;`); err != nil {
		return fmt.Errorf("failed to update an account: %w", err)
	}

	if err := tx2.Commit(); err != nil {
		return fmt.Errorf("failed to commit tx2: %w", err)
	}

	if err := pg.PrintAccounts(tx1, `SELECT * FROM accounts WHERE balance > 50;`); err != nil {
		return fmt.Errorf("failed to print accounts: %w", err)
	}

	if err := tx1.Commit(); err != nil {
		return fmt.Errorf("failed to commit tx1: %w", err)
	}

	return nil
}
