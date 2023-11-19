package configs

type Config struct {
	POSTGRESQL_URI string `json:"POSTGRESQL_URI" yaml:"POSTGRESQL_URI"`
	ServerAddr     string `json:"ServerAddr" yaml:"ServerAddr"`
	ServerPort     string `json:"ServerPort" yaml:"ServerPort"`
	RedisUri       string `json:"RedisUri" yaml:"RedisUri"`
}
