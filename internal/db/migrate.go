package db

import (
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (d *Database) MigrateDB(migrationsSource string) error {

	log.Println("running migrations....")

	driver, err := postgres.WithInstance(d.Client, &postgres.Config{})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsSource, "postgres", driver)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if err = m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			log.Println(err.Error())
			return err
		} else {
			log.Println("no changes")
		}
	}

	log.Println("migrations finished")

	return nil
}
