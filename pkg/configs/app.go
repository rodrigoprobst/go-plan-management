package configs

import (
	"log"

	"github.com/joho/godotenv"
)

const (
	Development = "development"
	Staging     = "stage"
	Production  = "production"
	Version     = "1.0.0"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func InitializeConfigs() {
	initializeApplicationConfigs()
	initializePostgresConfig()
}

type AppConfig struct {
	Env                string
	LogLevel           string
	AppName            string
	AppVersion         string
	Port               string
	CorsTrustedOrigins []string
}

type postgresConfig struct {
	Dsn string
}

var (
	ApplicationCfg *AppConfig
	PostgresCfg    *postgresConfig
)

func initializeApplicationConfigs() {
	if ApplicationCfg == nil {
		ApplicationCfg = &AppConfig{
			Env:                GetEnv("ENV", Production),
			LogLevel:           GetEnv("LOG_LEVEL", "info"),
			AppName:            GetEnv("APP_NAME", ""),
			AppVersion:         GetEnv("APP_VERSION", Version),
			Port:               GetEnv("HTTP_PORT", "80"),
			CorsTrustedOrigins: getEnvAsSlice("CORS_TRUSTED_ORIGINS", []string{}, " "),
		}
	}
}

func initializePostgresConfig() {
	if PostgresCfg == nil {
		PostgresCfg = &postgresConfig{
			Dsn: GetEnv("POSTGRES_DSN", ""),
		}
	}
}
