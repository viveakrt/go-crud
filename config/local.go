package config

func getLocalConfig() Configuration {
	return Configuration{
		AppPort:       8000,
		GrowthBookURL: "http://localhost:3100/api/features/key_prod_8dd8449251ae2a29",
		Sql: SQL{
			Host:     "localhost",
			Port:     "3306",
			UserName: "root",
			Password: "12345",
			Database: "netflix",
		},
	}
}
