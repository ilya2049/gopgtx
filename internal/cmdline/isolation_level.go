package cmdline

import (
	"database/sql"
	"errors"
	"os"

	"gopgtx/internal/pg"
)

var ErrSpecifyIsolationLevel = errors.New("specify transaction isolation level")

func ReadIsolationLevel() (sql.IsolationLevel, error) {
	if len(os.Args) == 1 {

		return 0, ErrSpecifyIsolationLevel
	}

	isolationLevel, err := pg.NewIsolationLevel(os.Args[1])
	if err != nil {
		return 0, err
	}

	return isolationLevel, nil
}
