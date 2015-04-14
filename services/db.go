package myapi

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUsername     = "gofast"
	dbDatabaseName = "gofast"
)

func GetDb() *sql.DB {
	dataSourceName := fmt.Sprintf("%s@/%s", dbUsername, dbDatabaseName)

	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to db %s: error: %v", dbDatabaseName, err))

	}

	if err := db.Ping(); err == nil {
		fmt.Println("Connected to db", dbDatabaseName)
	}

	return db
}

func InitDb() *sql.DB {
	db := GetDb()

	_, err := db.Exec(
		`CREATE TABLE BOOK (
			ID int NOT NULL AUTO_INCREMENT,
			NAME varchar(255) NOT NULL,
			AUTHOR varchar(255) NOT NULL,
			PRIMARY KEY (ID)
		) `,
	)

	if err != nil {
		fmt.Println("Failed to create the BOOK table:", err)
	}

	return db
}
