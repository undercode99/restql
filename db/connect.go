package db

import (
	"database/sql"
	"errors"
)

type Database struct {
	// contains filtered or unexported fields
	Username string
	Password string
	Host     string
	Port     string
	Database string
	SSLMode  string

	URL    string
	Driver string
	DB     *sql.DB
}

func NewDatabaseConnect(db *Database) (*Database, error) {
	return db.Connect()
}

func (d *Database) Connect() (*Database, error) {
	if d.Driver == "postgres" {
		if d.Port == "" {
			d.Port = "5432"
		}

		if d.SSLMode == "" {
			d.SSLMode = "disable"
		}
		connDsn := "user=" + d.Username + " password=" + d.Password + " host=" + d.Host + " port=" + d.Port + " dbname=" + d.Database + " sslmode=" + d.SSLMode
		// Connect to database
		con, err := sql.Open(d.Driver, connDsn)
		if err != nil {
			return nil, err
		}
		d.DB = con
		return d, nil
	}
	return nil, errors.New("driver not supported")
}

func (d *Database) Close() error {
	return d.DB.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.DB
}
