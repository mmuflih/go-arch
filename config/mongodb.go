package config

import (
	"github.com/globalsign/mgo"
	"github.com/mmuflih/envgo/conf"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-01-31 09:25
**/

func NewMongoDB(cfg conf.Config) (error, *mgo.Database) {
	address := cfg.GetString(`mongodb.address`)
	user := cfg.GetString(`mongodb.user`)
	pass := cfg.GetString(`mongodb.pass`)
	database := cfg.GetString(`mongodb.database`)
	port := cfg.GetString(`mongodb.port`)
	session, err := mgo.Dial(address + ":" + port)
	dbSession := session.DB(database)
	auth := cfg.GetBool(`mongodb.auth`)
	if auth {
		dbSession.Login(user, pass)
	}
	return err, dbSession
}
