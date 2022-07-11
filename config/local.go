package config

func getLocalConfig() Configuration {
	return Configuration{
		AppPort:       8000,
		GrowthBookURL: "http://localhost:3100/api/features/key_prod_6c6afc1e779450e4",
		Sql: SQL{
			Host:     "127.0.0.1",
			Port:     "3306",
			UserName: "root",
			Password: "12345",
			Database: "netflix",
		},
	}
}
