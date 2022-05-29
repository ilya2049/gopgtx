package pg

import (
	"context"
	"database/sql"
	"fmt"

	"gopgtx/internal/models"
)

func PrintAccounts(ctx context.Context, tx *sql.Tx) error {
	rows, err := tx.QueryContext(ctx, `SELECT * FROM accounts WHERE balance > 50;`)
	if err != nil {
		return err
	}

	var (
		account models.Account
	)

	fmt.Println("----------- Accounts -----------")

	for rows.Next() {
		if err := rows.Scan(&account.ID, &account.Balance); err != nil {
			return err
		}

		fmt.Println(account)
	}

	return nil
}

func UpdateAccount(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`UPDATE accounts SET balance = 51.0 WHERE balance < 50.0;`)

	return err
}

func DeleteAccount(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`DELETE FROM accounts WHERE balance < 50.0;`)

	return err
}
