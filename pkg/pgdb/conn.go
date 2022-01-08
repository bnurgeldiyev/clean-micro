package pgdb

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func NewPgConn(dbConn string, dbMaxConn int) (conn *sqlx.DB, err error) {

	conn, err = sqlx.Connect("postgres", dbConn)
	if err != nil {
		eMsg := "error creating connection sqlx"
		log.WithError(err).Error(eMsg)
		//err = status.Error(codes.Code(500), eMsg)
		return
	}

	conn.SetMaxOpenConns(dbMaxConn)

	return
}
