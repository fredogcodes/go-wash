package main

import (
	"fmt"
	"go_carwash/config"
	"go_carwash/database"
	"go_carwash/store"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	config := config.LoadConfig(".")

	db := database.NewDBConn(config)

	store := store.NewStore(db)
	fmt.Println("Database connection successful", store)

	fmt.Println("Running migrations")
	runDBMigration(config.MigrationPath, config.DBSOURCE)
	fmt.Println("Migrations complete")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
	fmt.Println("Server started on port 8080")

}

func runDBMigration(migrationFile, DBSource string) {

	migration, err := migrate.New(migrationFile, DBSource)
	if err != nil {
		fmt.Println("Error creating migration: ", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println("Error running migration: ", err)
	}

}
