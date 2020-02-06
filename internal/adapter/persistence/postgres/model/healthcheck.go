package model

import (
	"log"

	"github.com/jasonsmithj/hchd/internal/adapter/configuration"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func postgresConnect() *gorm.DB {
	db, err := gorm.Open("postgres", configuration.GetPostgresUri())
	if err != nil {
		log.Printf("db connect failed %s", err)
	}
	return db
}

func PostgresHealthcheck() error {
	c := postgresConnect()
	err := c.DB().Ping()

	defer c.Close()

	return err
}
