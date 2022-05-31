package scenarios

import (
	"database/sql"
	"errors"
	"fmt"
)

type ScenarioFunc = func(*sql.DB, sql.IsolationLevel) error

type Scenario struct {
	ID   int
	Name string
	Func ScenarioFunc
}

func (s Scenario) String() string {
	return fmt.Sprintf("[%d] %s", s.ID, s.Name)
}

var Scenarios = []Scenario{
	{
		ID:   1,
		Name: "Select committed.",
		Func: selectCommitted,
	},
	{
		ID:   2,
		Name: "Update deleted.",
		Func: updateDeleted,
	},
	{
		ID:   3,
		Name: "Update deleted. Waiting for the transaction to complete.",
		Func: updateDeletedWaitingForTxComplete,
	},
}

var ErrScenarioNotFound = errors.New("scenario with this id not found")

func Get(scenarioIndex int) (Scenario, error) {
	scenarioIndex = scenarioIndex - 1

	if scenarioIndex < 0 || scenarioIndex >= len(Scenarios) {
		return Scenario{}, ErrScenarioNotFound
	}

	return Scenarios[scenarioIndex], nil
}
