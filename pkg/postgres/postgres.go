package postgres

import (
	"database/sql"
	"fmt"
)

type ssl string

const (
	sslEnabled  ssl = "enable"
	sslDisabled ssl = "disable"
)

type Postgres struct {
	dsn    string
	driver string
	db     *sql.DB
}

func New(host string, port int, username string, password string, database string, SSLEnabled bool, driver string) *Postgres {
	dsn := generateDsn(host, port, username, password, database, SSLEnabled)
	return &Postgres{
		dsn:    dsn,
		driver: driver,
	}
}

func generateDsn(host string, port int, username string, password string, database string, SSLEnabled bool) string {
	sslValue := sslDisabled
	if SSLEnabled {
		sslValue = sslEnabled
	}

	result := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		username,
		password,
		database,
		sslValue,
	)
	return result
}

func (p *Postgres) Connect() error {
	db, err := sql.Open(p.driver, p.dsn)
	if err != nil {
		return err
	}

	p.db = db
	return nil
}

func (p *Postgres) Disconnect() error {
	return p.db.Close()
}

func (p *Postgres) RunQuery(query string, args ...any) (*sql.Rows, error) {
	rows, err := p.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (p *Postgres) DB() *sql.DB {
	return p.db
}

type DBConnection struct {
	Postgres *Postgres
}

func (db *DBConnection) GetDBConnection() (*sql.DB, error) {
	if err := db.Postgres.Connect(); err != nil {
		return nil, err
	}

	if _, err := db.Postgres.RunQuery("SELECT 1+1;"); err != nil {
		return nil, err
	}

	return db.Postgres.db, nil
}
