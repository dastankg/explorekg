package db

import (
	"explorekg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *config.Config) (*Db, error) {
	db, err := gorm.Open(postgres.Open(conf.Db.DATABASE_URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Db{db}, nil
}
