package composites

import (
	"clean-micro/pkg/pgdb"
	"fmt"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type PostgresComposite struct {
	DB *sqlx.DB
}

func NewPostgresComposite(dbConn string, dbMaxConn int) (*PostgresComposite, error) {

	fmt.Println(dbConn)

	conn, err := pgdb.NewPgConn(dbConn, dbMaxConn)
	if err != nil {
		eMsg := "error in pgdb.NewPgConn"
		log.WithError(err).Error(eMsg)
		return nil, err
	}

	return &PostgresComposite{
		DB: conn,
	}, nil
}
