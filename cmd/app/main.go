package main

import (
	"context"
	"log"

	"gopgtx/internal/cmdline"
	"gopgtx/internal/pg"
	"gopgtx/internal/pg/scenarios"
)

func main() {
	isolationLevel, err := cmdline.ReadIsolationLevel()
	if err != nil {
		log.Println(err)

		return
	}

	cmdline.PrintScenarios()
	scenarioID, err := cmdline.ReadScenarioID()
	if err != nil {
		log.Println(err)

		return
	}

	scenario, err := scenarios.Get(scenarioID)
	if err != nil {
		log.Println(err)

		return
	}

	db, closeDB, err := pg.NewConnection(pg.DefaultConfig())
	if err != nil {
		log.Println(err)

		return
	}

	defer closeDB()

	ctx := context.Background()

	if err := scenario.Func(ctx, db, isolationLevel); err != nil {
		log.Println("scenario failed: " + err.Error())

		return
	}

	log.Println("OK")
}
