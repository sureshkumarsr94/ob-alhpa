package migrations

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	cfg "infopack.co.in/offybox/app/configs"
	"log"
)

func RunMigrations() error {
	db, err := sql.Open("mysql", cfg.GetConfig().Mysql.GetMysqlConnectionForMigrate())
	if err != nil {
		return fmt.Errorf("could not open db connection: %v", err)
	}
	defer db.Close()

	if err := goose.SetDialect("mysql"); err != nil {
		return fmt.Errorf("could not set goose dialect: %v", err)
	}

	// Apply product-level migrations
	if err := applyMigrations(db, "./migrations/schema"); err != nil {
		return fmt.Errorf("could not apply product migrations: %v", err)
	}

	log.Println("Migrations applied successfully")
	return nil
}

func applyMigrations(db *sql.DB, migrationPath string) error {
	if err := goose.Up(db, migrationPath); err != nil {
		return fmt.Errorf("could not apply migrations from %s: %v", migrationPath, err)
	}
	return nil
}
