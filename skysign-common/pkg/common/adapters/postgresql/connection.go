package postgresql

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsnTemplate = "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s"

var (
	host     = "localhost"
	user     = "tomofiles"
	password = "pc+tomofiles"
	dbname   = "##dbname##"
	port     = "5432"
	sslmode  = "disable"
)

// NewPostgresqlConnection .
func NewPostgresqlConnection(paramDbname string) (*gorm.DB, error) {
	if envHost := os.Getenv("DB_HOST"); envHost != "" {
		host = envHost
	}
	if envUser := os.Getenv("DB_USERNAME"); envUser != "" {
		user = envUser
	}
	if envPassword := os.Getenv("DB_PASSWORD"); envPassword != "" {
		password = envPassword
	}
	if envPort := os.Getenv("DB_PORT"); envPort != "" {
		port = envPort
	}
	if envSslmode := os.Getenv("DB_SSL_ENABLED"); envSslmode != "" {
		sslmode = envSslmode
	}

	dbname = paramDbname

	dsn := fmt.Sprintf(dsnTemplate, host, user, password, dbname, port, sslmode)

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to Postgresql")
	}

	return db, nil
}
