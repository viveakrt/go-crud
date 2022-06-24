package config

type Configuration struct {
	AppPort       int
	GrowthBookURL string
	Sql           SQL
}

type SQL struct {
	Port     string
	UserName string
	Password string
	Host     string
	Database string
}
