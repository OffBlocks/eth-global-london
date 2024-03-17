package db

import (
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	p "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(p.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	instance, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}

	driver, err := postgres.WithInstance(instance, &postgres.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalln(err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalln(err)
	}

	return Handler{db}
}
