package repository

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func RunMigrations(db *Database, direction string) error {
	migrationsDir := "migrations"

	// Read all SQL files in the migrations directory
	files, err := filepath.Glob(filepath.Join(migrationsDir, "*.sql"))
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %v", err)
	}

	// Sort files by name to ensure correct order
	for _, file := range files {
		fmt.Printf("Applying migration: %s\n", file)

		// Read the migration file
		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %v", file, err)
		}

		// Split the content into individual statements
		statements := strings.Split(string(content), ";")

		// Execute each statement
		for _, stmt := range statements {
			stmt = strings.TrimSpace(stmt)
			if stmt == "" {
				continue
			}

			if err := db.DB.Exec(stmt).Error; err != nil {
				return fmt.Errorf("failed to execute statement in %s: %v", file, err)
			}
		}
	}

	return nil
}
