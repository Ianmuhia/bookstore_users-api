package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUsersUsername = "mysqlUsersUsername"
	mysqlUsersSchema   = "mysqlUsersSchema"
	mysqlUsersHost     = "mysqlUsersHost"
	mysqlUsersPassword = "mysqlUsersPassword"
)

var (
	Client *sql.DB

	username = os.Getenv(mysqlUsersUsername)
	host     = os.Getenv(mysqlUsersHost)
	schema   = os.Getenv(mysqlUsersSchema)
	password = os.Getenv(mysqlUsersPassword)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")

}
