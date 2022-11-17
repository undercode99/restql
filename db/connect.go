package db

import (
	"database/sql"

	_ "github.com/lib/pq" // add this

	"errors"
)

type Database struct {
	// contains filtered or unexported fields
	Name     string  `json:"name"` // Name connection of database
	Username string  `json:"username"`
	Password string  `json:"password"`
	Host     string  `json:"host"`
	Port     string  `json:"port"`
	Database string  `json:"database"`
	SSLMode  string  `json:"sslmode"`
	Dsn      string  `json:"dsn"`
	Driver   string  `json:"driver"`
	DB       *sql.DB `json:"-"`
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
		err = con.Ping()
		if err != nil {
			return nil, err
		}

		d.Dsn = connDsn
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

type ListDatabaseConnect struct {
	// contains filtered or unexported fields
	Databases map[string]*Database
}

func NewListDatabaseConnect() *ListDatabaseConnect {
	return &ListDatabaseConnect{}
}

func (l *ListDatabaseConnect) AddConnection(name string, db *Database) error {
	if l.Databases == nil {
		l.Databases = make(map[string]*Database)
	}
	l.Databases[name] = db
	return nil
}

func (l *ListDatabaseConnect) GetConnection(name string) (*Database, error) {
	if l.Databases == nil {
		return nil, errors.New("database not found")
	}
	if _, ok := l.Databases[name]; ok {
		return l.Databases[name], nil
	}
	return nil, errors.New("database not found")
}

func (l *ListDatabaseConnect) CheckExist(name string) bool {
	if l.Databases == nil {
		return false
	}
	if _, ok := l.Databases[name]; ok {
		return true
	}
	return false
}

func (l *ListDatabaseConnect) GetList() map[string]*Database {
	return l.Databases
}

func (l *ListDatabaseConnect) Close() error {
	for _, db := range l.Databases {
		err := db.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
