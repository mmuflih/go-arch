package config

import (
	"github.com/mmuflih/envgo/conf"
	"github.com/mmuflih/go-di-arch/app"
	"gopkg.in/mgo.v2"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-01-31 09:25
**/

type MongoConnections struct {
	Conn1 *mgo.Database
}

func NewMongoDB2(cfg conf.Config) (*MongoConnections, error) {
	address := cfg.GetString(`mongodb.address`)
	user := cfg.GetString(`mongodb.user`)
	pass := cfg.GetString(`mongodb.pass`)
	database := cfg.GetString(`mongodb.database`)
	port := cfg.GetString(`mongodb.port`)
	auth := cfg.GetBool(`mongodb.auth`)
	app.Logger("=>>", address, port, database)

	session, err := mgo.Dial(address + ":" + port)
	if err != nil {
		app.Logger("<>>", "Mongodb Conn", err)
		return nil, err
	}
	dbSession1 := session.DB(database)
	app.Logger("=>>", "Mongo auth", auth)
	if auth {
		err := dbSession1.Login(user, pass)
		if err != nil {
			app.Logger("<>>", "Mongodb Conn", err)
			return nil, err
		}
	}
	return &MongoConnections{
		Conn1: dbSession1,
	}, nil
}
