package initialize

import (
	"myapp/config"
	"myapp/internal/http"
	"myapp/pkg/db/mysql"
)

func InitializeApp() {

	config.InitConfig()
	mysql.CreateClient()
	http.InitHTTPServer()
}
