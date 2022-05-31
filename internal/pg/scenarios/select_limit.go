package scenarios

import (
	"context"
	"database/sql"
	"fmt"

	"gopgtx/internal/models"
	"gopgtx/internal/pg"
)

func selectLimitWithInserting(db *sql.DB, _ sql.IsolationLevel) error {
	tx1, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return fmt.Errorf("failed to open tx1: %w", err)
	}

	tx2, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return fmt.Errorf("failed to open tx2: %w", err)
	}

	if err := pg.PrintAccounts(tx1,
		`SELECT * FROM accounts ORDER BY balance LIMIT 2 FOR UPDATE;`,
	); err != nil {
		return fmt.Errorf("failed to print accounts: %w", err)
	}

	if err := pg.InsertAccount(tx2, models.Account{Balance: 10}); err != nil {
		return fmt.Errorf("failed to insert an account: %w", err)
	}

	if err := tx2.Commit(); err != nil {
		return fmt.Errorf("failed to commit tx2: %w", err)
	}

	tx3, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return fmt.Errorf("failed to open tx3: %w", err)
	}

	if err := pg.PrintAccounts(tx3,
		`SELECT * FROM accounts ORDER BY balance LIMIT 1 FOR UPDATE;`,
	); err != nil {
		return fmt.Errorf("failed to print accounts: %w", err)
	}

	if err := tx1.Commit(); err != nil {
		return fmt.Errorf("failed to commit tx1: %w", err)
	}

	if err := tx3.Commit(); err != nil {
		return fmt.Errorf("failed to commit tx3: %w", err)
	}

	return nil
}

func selectLimitWaitingForTxComplete(db *sql.DB, isolationLevel sql.IsolationLevel) error {
	tx1, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: isolationLevel})
	if err != nil {
		return fmt.Errorf("failed to open tx1: %w", err)
	}

	tx2, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: isolationLevel})
	if err != nil {
		return fmt.Errorf("failed to open tx2: %w", err)
	}

	if err := pg.PrintAccounts(tx1,
		`SELECT * FROM accounts ORDER BY balance LIMIT 1 FOR UPDATE;`,
	); err != nil {
		return fmt.Errorf("failed to print accounts: %w", err)
	}

	if err := pg.PrintAccounts(tx2,
		`SELECT * FROM accounts ORDER BY balance LIMIT 1 FOR UPDATE;`,
	); err != nil {
		return fmt.Errorf("failed to print accounts: %w", err)
	}

	// deadlock

	return nil
}
