package bookstore_users_db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlBookstoreUsersUsername = "MYSQL_BOOKSTORE_USERS_DB_USERNAME"
	mysqlBookstoreUsersPassword = "MYSQL_BOOKSTORE_USERS_DB_PASSWORD"
	mysqlBookstoreUsersHost     = "MYSQL_BOOKSTORE_USERS_DB_HOST"
	mysqlBookstoreUsersPort     = "MYSQL_BOOKSTORE_USERS_DB_PORT"
	mysqlBookstoreUsersDatabase = "MYSQL_BOOKSTORE_USERS_DB_DATABASE"
)

var (
	Client *sql.DB

	username = os.Getenv(mysqlBookstoreUsersUsername)
	password = os.Getenv(mysqlBookstoreUsersPassword)
	host     = os.Getenv(mysqlBookstoreUsersHost)
	port     = os.Getenv(mysqlBookstoreUsersPort)
	database = os.Getenv(mysqlBookstoreUsersDatabase)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		username, password, host, port, database)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("database successfully configured")
}
