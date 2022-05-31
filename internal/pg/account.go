package pg

import (
	"database/sql"
	"fmt"

	"gopgtx/internal/models"
)

func PrintAccounts(tx *sql.Tx, query string) error {
	rows, err := tx.Query(query)
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

func UpdateAccount(tx *sql.Tx, query string) error {
	_, err := tx.Exec(query)

	return err
}

func DeleteAccount(tx *sql.Tx, query string) error {
	_, err := tx.Exec(query)

	return err
}
