package dao

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"log"
	"time"
	"trace_demo/configs"
)

var dbMap map[string]*gorm.DB

type Config struct {
	DSNList []configs.MysqlDSN
}

func Startup(config Config) {
	dbMap = make(map[string]*gorm.DB)
	for _, item := range config.DSNList {
		db, err := gorm.Open(mysql.Open(item.DSN), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		dbMap[item.Name] = db

		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal(err)
		}

		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(time.Hour)
	}
}

func GetDB(ctx context.Context, name string) *gorm.DB {
	return dbMap[name].Session(&gorm.Session{
		Context: ctx,
		Logger:  NewLogger(ctx, logger2.Info),
	})
}

type Dao struct {
	db *gorm.DB
}

func New(ctx context.Context, name string) *Dao {
	return &Dao{db: GetDB(ctx, name)}
}
