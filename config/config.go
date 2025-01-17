package config

import (
	"fmt"
	"os"
)

var curConfig *Config

func init() {
	curConfig = Load()
}

func GetConfig() *Config {
	return curConfig
}

type Config struct {
	MySQLDSN   string
	DorisDSN   string
	ServerPort string
}

// Load 从环境变量或默认值加载配置
func Load() *Config {
	cfg := &Config{}
	cfg.MySQLDSN = getEnv("MYSQL_DSN", "root:password@tcp(127.0.0.1:3306)/dsp_db?charset=utf8mb4&parseTime=True&loc=Local")
	cfg.DorisDSN = getEnv("DORIS_DSN", "root:@tcp(127.0.0.1:9030)/doris_db")
	cfg.ServerPort = getEnv("SERVER_PORT", ":8080")

	fmt.Printf("[CONFIG] Loaded config: %+v\n", cfg)
	return cfg
}

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}
