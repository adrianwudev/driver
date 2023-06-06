package migration

import (
	"driver/config"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migration struct {
	client *migrate.Migrate
}

var dbConfig = config.GetConfig()

var dbUrl = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
	dbConfig.User,
	dbConfig.Password,
	dbConfig.Host,
	dbConfig.Port,
	dbConfig.DBName,
)

func NewMigration() *Migration {
	m := Migration{}
	path := "file://../assets/db/migration"
	var err error
	m.client, err = migrate.New(path, dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	return &m
}

func (m *Migration) To(targetVersion uint) {
	if err := m.client.Migrate(targetVersion); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	afterVersion, _, _ := m.client.Version()
	fmt.Printf("Migration to version:%d success", afterVersion)
}

func (m *Migration) Up() {
	if err := m.client.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	afterVersion, _, _ := m.client.Version()
	fmt.Printf("Migration up version:%d success", afterVersion)
}
