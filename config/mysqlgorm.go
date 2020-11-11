package config

import (
	"log"
	"os"
	"time"

	"github.com/mmuflih/envgo/conf"
	"github.com/mmuflih/go-di-arch/app"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MyConn struct {
	Conn1 *gorm.DB
}

func NewMysqlGormConn(cfg conf.Config) (*MyConn, error) {
	dbUser := cfg.GetString(`mysql.user`)
	dbPass := cfg.GetString(`mysql.pass`)
	dbName := cfg.GetString(`mysql.database`)
	dbHost := cfg.GetString(`mysql.address`)
	dbPort := cfg.GetString(`mysql.port`)
	query := cfg.GetString(`mysql.query`)

	if query == "" {
		query = "parseTime=true&loc=Asia%2FJakarta"
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Warn,
			Colorful:      true,
		},
	)

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?" + query
	app.Logger("=> GORM DSN ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		app.Logger("<>", "GORM Conn", err)
		return nil, err
	}

	return &MyConn{
		Conn1: db,
	}, nil
}
