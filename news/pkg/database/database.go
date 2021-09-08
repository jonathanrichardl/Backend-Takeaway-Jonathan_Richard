package database

import (
	"database/sql"
	"errors"
	"fmt"
	mysql "news/pkg/database/mysql"
	"strings"
)

type DBInstance struct {
	db *sql.DB
}

type RetrievedData struct {
	Data *sql.Rows
}

func (a *DBInstance) AddData(Query string) error {
	_, err := a.db.Query(Query)
	return err

}

func (a *DBInstance) RetrieveData(Query string) (*RetrievedData, error) {
	rows, err := a.db.Query(Query)
	if err != nil {
		return nil, err
	}
	return &RetrievedData{rows}, nil
}

func NewDatabase(DatabaseManagementSystem string, Username string, Password string, Address string, DatabaseName string) (*DBInstance, error) {
	var DB *sql.DB
	var err error
	switch strings.ToLower(DatabaseManagementSystem) {
	case "mysql":
		DB, err = mysql.Connect(fmt.Sprintf("%s:%s@tcp(%s)/%s", Username, Password, Address, DatabaseName))
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New(fmt.Sprintf("%s database is not implemented", DatabaseManagementSystem))
	}
	return &DBInstance{DB}, nil
}
