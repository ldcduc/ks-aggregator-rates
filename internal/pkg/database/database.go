package database

import (
	_ "database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDb() *gorm.DB {
	dsn := "host=localhost port=54320 user=postgres password=my_password dbname=ks_aggregator_rates sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// db, err := gorm.Open("postgres", "host=localhost port=54320 user=postgres dbname=ks_aggregator_rates sslmode=disable password=my_password")
	if err != nil {
		panic("Failed to connect to database")
	}

	return db
}

func CloseDb(db *gorm.DB) {
	sql_db, err := db.DB()

	if err != nil {
		panic("Failed to close to database")
	}

	sql_db.Close()
}
