package config

type Config struct {
	DBConfig     DBConfig
	MaxRetiresDB int
	SecretKey    []byte
}

type DBConfig struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

var CurrentConfig = Config{
	DBConfig: DBConfig{
		Driver:   "mysql",
		Username: "root",
		Password: "password",
		Host:     "db",
		Port:     3306,
		Database: "vote",
	},
	SecretKey:    []byte("MTL_SECRET_KEY"),
	MaxRetiresDB: 3,
}
