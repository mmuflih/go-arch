package config

import (
	"log"
	"os"
	"time"

	"github.com/mmuflih/envgo/conf"
	"github.com/mmuflih/go-di-arch/app"
	"gopkg.in/mgo.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-01-31 09:25
**/

type MongoDB struct {
	MainDB *mgo.Database
}

type MySQL struct {
	MainDB *gorm.DB
}

func NewMySQLConnections(cfg conf.Config) (*MySQL, error) {
	app.Logger("Initial Mysql Connection")
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
	app.Logger("GORM DSN ", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		app.Error(err)
		return nil, err
	}

	if cfg.GetString("env") == "dev" || cfg.GetString("env") == "staging" {
		db = db.Debug()
	}
	app.Success("Mysql Connected")
	return &MySQL{
		MainDB: db,
	}, nil
}

func NewMongoDBConnections(cfg conf.Config) (*MongoDB, error) {
	app.Logger("Initial mongodb")
	address := cfg.GetString(`mongodb.address`)
	user := cfg.GetString(`mongodb.user`)
	pass := cfg.GetString(`mongodb.pass`)
	database := cfg.GetString(`mongodb.database`)
	port := cfg.GetString(`mongodb.port`)
	auth := cfg.GetBool(`mongodb.auth`)
	app.Logger(address, port, database)

	session, err := mgo.Dial(address + ":" + port)
	if err != nil {
		app.Error(err)
		return nil, err
	}
	mongoDBSession := session.DB(database)
	app.Logger("Mongo auth", auth)
	if auth {
		err := mongoDBSession.Login(user, pass)
		if err != nil {
			app.Error(err)
			return nil, err
		}
	}
	app.Success("mongodb connected")
	return &MongoDB{
		MainDB: mongoDBSession,
	}, nil
}
