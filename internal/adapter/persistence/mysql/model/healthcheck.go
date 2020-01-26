package model

import (
	"log"

	"github.com/jasonsmithj/healthcheck-container-http-mysql/internal/adapter/configuration"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func connect() *gorm.DB {
	db, err := gorm.Open("mysql", configuration.GetDBUri())
	if err != nil {
		log.Printf("db connect failed %s", err)
	}
	return db
}

func Healthcheck() error {
	c := connect()
	err := c.DB().Ping()

	defer c.Close()

	return err
}
