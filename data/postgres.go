package data

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"shop-backend/config"
	"shop-backend/utils/logger"
)

var db *sqlx.DB

func init() {
	logger.L.Infoln("connecting to postgresql")
	var err error
	db, err = sqlx.Connect("postgres", config.Cfg.PostgresUrl)
	if err != nil {
		logger.L.Fatalln(err, "postgresql connection failed")
	}
	err = db.Ping() //ping to make sure we are connected
	if err != nil {
		logger.L.Fatalln(err, "postgresql connection failed")
	}
	db.SetMaxOpenConns(200)
	logger.L.Infoln("postgresql database connected")

}

func DB() *sqlx.DB {
	return db
}
