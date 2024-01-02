package config

type Config struct {
	DBConfig     DBConfig
	MaxRetiresDB int
	SecretKey    []byte
	RedisConfig  RedisConfig
}

type RedisConfig struct {
	Address  string
	Password string
	DB       int
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
	SecretKey: []byte("MTL_SECRET_KEY"),
	RedisConfig: RedisConfig{
		Address:  "redis:6379",
		Password: "",
		DB:       0,
	},
	MaxRetiresDB: 3,
}
