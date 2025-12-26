package main

import (
	"fmt"
	"log"

	database "github.com/faizallmaullana/lenteng-agung/backend/db/db_connection"
	"github.com/faizallmaullana/lenteng-agung/backend/env"
	"github.com/faizallmaullana/lenteng-agung/backend/internal/router"
)

func main() {
	// initialize DB provider
	provider, err := database.NewGormProviderFromEnv()
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	defer provider.Close()

	// build router with wired handlers
	r, err := router.NewRouter(provider)
	if err != nil {
		log.Fatal("failed to create router: ", err)
	}

	config := env.Get()
	if err := r.Run(fmt.Sprintf(":%s", config.App.Port)); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
