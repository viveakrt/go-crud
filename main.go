package main

import (
	"myapp/internal/initialize"
	"myapp/pkg/db/mysql"
)

func main() {
	initialize.InitializeApp()
	mysql.CreateClient()
}
