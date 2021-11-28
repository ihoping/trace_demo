package dao

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
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
		db, err := gorm.Open(mysql.Open(item.DSN), &gorm.Config{
			Logger: &sqlLogger{os.Stdout},
		})
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
		Logger:  &sqlLogger{os.Stdout},
	})
}

type sqlLogger struct {
	io.Writer
}

func (l sqlLogger) LogMode(level logger.LogLevel) logger.Interface {
	fmt.Println(level)
	return &l
}

func (l sqlLogger) Info(ctx context.Context, msg string, others ...interface{}) {
	l.Write([]byte(msg))
}

func (l sqlLogger) Warn(ctx context.Context, msg string, others ...interface{}) {
	l.Write([]byte(msg))
}

func (l sqlLogger) Error(ctx context.Context, msg string, others ...interface{}) {
	l.Write([]byte(msg))
}

func (l sqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	l.Write([]byte(sql + ctx.Value("request_id").(string)))
}
