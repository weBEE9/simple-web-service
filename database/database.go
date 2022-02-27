package database

import (
	"fmt"
	"weBEE9/simple-web-service/config"

	_ "github.com/lib/pq"

	"xorm.io/core"
	"xorm.io/xorm"
)

func NewEngine(cfg config.Config) (*xorm.Engine, error) {
	e, err := xorm.NewEngine(cfg.DB.Driver, getPostgresConnetionInfo(cfg))
	if err != nil {
		return nil, err
	}

	e.SetMapper(core.GonicMapper{})

	if err := e.Sync2(tables...); err != nil {
		return nil, err
	}
	e.ShowSQL(cfg.DB.Debug)

	return e, nil
}

func getPostgresConnetionInfo(cfg config.Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Database)
}
