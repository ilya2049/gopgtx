package cmdline

import (
	"fmt"
	"gopgtx/internal/pg/scenarios"
)

func PrintScenarios() {
	fmt.Println("------------- Test scenarios -------------")
	fmt.Println("Chose scenario id:")

	for _, scenario := range scenarios.Scenarios {
		fmt.Println(scenario)
	}
}

func ReadScenarioID() (int, error) {
	var scenarioID int

	_, err := fmt.Scan(&scenarioID)
	if err != nil {
		return 0, fmt.Errorf("failed to read scenario id: %w", err)
	}

	return scenarioID, nil
}
