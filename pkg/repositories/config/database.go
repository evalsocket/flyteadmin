package config

import "time"

// Database config. Contains values necessary to open a database connection.
type DbConfig struct {
	BaseConfig
	Host                  string        `json:"host"`
	Port                  int           `json:"port"`
	DbName                string        `json:"dbname"`
	User                  string        `json:"user"`
	Password              string        `json:"password"`
	ExtraOptions          string        `json:"options"`
	MaxOpenConnection     int           `json:"maxOpenConnection"`
	MaxIdleConnection     int           `json:"maxIdleConnection"`
	ConnectionMaxLifetime time.Duration `json:"connectionMaxLifetime"`
}
