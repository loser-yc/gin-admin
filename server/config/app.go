package config


type Config struct {
	AppName string
	Mysql MySQLConfig
}

type MySQLConfig struct {
	Host string
	User string
	Password string
	Dbname string
	Config string
}