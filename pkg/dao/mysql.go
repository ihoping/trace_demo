package dao

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"time"
)

var dbMap map[string]*gorm.DB

type Config struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func Startup(configs []Config) error {
	dbMap = make(map[string]*gorm.DB)
	for _, item := range configs {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", item.User, item.Password, item.Host, item.Port, item.Database)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: &logger{LogLevel: logger2.Info},
		})
		if err != nil {
			return err
		}
		dbMap[item.Name] = db

		sqlDB, err := db.DB()
		if err != nil {
			return err
		}

		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(time.Hour)
	}
	return nil
}

func GetDB(name string, ctx context.Context) (*gorm.DB, error) {
	if db, ok := dbMap[name]; ok {
		return db.Session(&gorm.Session{
			Context: ctx,
		}), nil
	}
	return nil, errors.New("db not found")
}

func Shutdown() error {
	for _, v := range dbMap {
		sqlDB, err := v.DB()
		if err != nil {
			return err
		}
		err = sqlDB.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

type Dao struct {
	db *gorm.DB
}

func New(dbName string, ctx context.Context) (*Dao, error) {
	db, err := GetDB(dbName, ctx)
	if err != nil {
		return nil, err
	}

	return &Dao{db: db}, nil
}
