package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewDBConnect(DSN string) (*sqlx.DB, error) {

	db, err := sqlx.Connect("postgres", DSN)
	//sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("ошибка подключения: ", err)
		return nil, err
	}
	//defer db.Close()
	fmt.Println("check ping")
	err = db.Ping()
	if err != nil {
		fmt.Println("ошибка ping: ", err)
		return nil, err
	}
	return db, nil
}
