package mysql

import (
	"database/sql"
	"fmt"
	"myapp/config"
	"os"
)

var DB *sql.DB

func getConnString(sqlConfig config.SQL) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", sqlConfig.UserName,
		sqlConfig.Password, sqlConfig.Host,
		sqlConfig.Port, sqlConfig.Database)
}

func CreateClient() {
	db, dbErr := sql.Open("mysql",
		getConnString(config.Config.Sql))
	if dbErr != nil {
		fmt.Println(dbErr)
		os.Exit(1)
	}
	DB = db

	pingErr := DB.Ping()
	if pingErr != nil {
		fmt.Println(dbErr)
		os.Exit(1)
	}

	prepareStatements()
}
