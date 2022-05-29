package scenarios

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type ScenarioFunc = func(context.Context, *sql.DB, sql.IsolationLevel) error

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
}

var ErrScenarioNotFound = errors.New("scenario with this id not found")

func Get(scenarioIndex int) (Scenario, error) {
	scenarioIndex = scenarioIndex - 1

	if scenarioIndex < 0 || scenarioIndex >= len(Scenarios) {
		return Scenario{}, ErrScenarioNotFound
	}

	return Scenarios[scenarioIndex], nil
}