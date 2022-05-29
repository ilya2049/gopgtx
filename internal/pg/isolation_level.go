package pg

import (
	"database/sql"
	"errors"
)

var ErrUnsupportedIsolationLevel = errors.New("unsupported isolation level")

func NewIsolationLevel(isolationLevelAsString string) (sql.IsolationLevel, error) {
	switch isolationLevelAsString {
	case "read_committed":
		return sql.LevelReadCommitted, nil
	case "repeatable_read":
		return sql.LevelRepeatableRead, nil
	case "serializable":
		return sql.LevelSerializable, nil
	}

	return 0, ErrUnsupportedIsolationLevel
}
