package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	portKey            = "PORT"

	postgresHostConfKey     = "PG_DB_HOST"
	postgresPortConfKey     = "PG_DB_PORT"
	postgresDBConfKey       = "PG_DB_NAME"
	postgresUserConfKey     = "PG_DB_USER"
	postgresPasswordConfKey = "PG_DB_PASS"

	postgresSSLModeConfKey     = "PG_SSL_MODE"
	postgresRootCertLocConfKey = "PG_ROOT_CERT_LOC"

	postgresMaxOpenConnsKey = "PG_MAX_OPEN_CONNS"
	postgresMaxIdleConnsKey = "PG_MAX_IDLE_CONNS"
	postgresMaxIdleTimeKey  = "PG_MAX_IDLE_TIME"

	jwtSecret = "JWT_SECRET"
)

type Config struct {
	Port        string

	JwtSecret string

	PostgresHost        string
	PostgresPort        string
	PostgresUser        string
	PostgresDB          string
	PostgresPassword    string
	PostgresSSLMode     string
	PostgresRootCertLoc string

	PostgresMaxOpenConns int
	PostgresMaxIdleConns int
	PostgresMaxIdleTime  time.Duration
}

var Conf *Config

func New() (*Config, error) {

	vars := &confVars{}

	port := vars.mandatory(portKey)

	postgresHost := vars.mandatory(postgresHostConfKey)
	postgresPort := vars.mandatory(postgresPortConfKey)
	postgresDB := vars.mandatory(postgresDBConfKey)
	postgresUser := vars.mandatory(postgresUserConfKey)
	postgresPassword := vars.mandatory(postgresPasswordConfKey)

	postgresSSLMode := vars.optional(postgresSSLModeConfKey, "disable")
	postgresRootCertLoc := vars.optional(postgresRootCertLocConfKey, "")

	postgresMaxOpenConns := vars.mandatoryInt(postgresMaxOpenConnsKey)
	postgresMaxIdleConns := vars.mandatoryInt(postgresMaxIdleConnsKey)
	postgresMaxIdleTime := vars.mandatoryDuration(postgresMaxIdleTimeKey)


	jwtSecretKey := vars.mandatory(jwtSecret)
	config := &Config{
		Port:                 port,
		PostgresHost:         postgresHost,
		PostgresPort:         postgresPort,
		PostgresDB:           postgresDB,
		PostgresUser:         postgresUser,
		PostgresPassword:     postgresPassword,
		PostgresSSLMode:      postgresSSLMode,
		PostgresRootCertLoc:  postgresRootCertLoc,
		PostgresMaxOpenConns: postgresMaxOpenConns,
		PostgresMaxIdleConns: postgresMaxIdleConns,
		PostgresMaxIdleTime:  postgresMaxIdleTime,
		JwtSecret:            jwtSecretKey,
	}

	if err := vars.Error(); err != nil {
		return nil, fmt.Errorf("Config: environment variables: %v", err)
	}

	Conf = config

	return config, nil
}

type confVars struct {
	missing   []string //name of the mandatory environment variable that are missing
	malformed []string //errors describing malformed environment varibale values
}

func (vars confVars) optional(key, fallback string) string {
	val := os.Getenv(key)

	if val == "" {
		return fallback
	}
	return val
}
func (vars *confVars) mandatory(key string) string {
	val := os.Getenv(key)

	if val == "" {
		vars.missing = append(vars.missing, key)
		return ""
	}

	return val
}

func (vars *confVars) mandatoryDuration(key string) time.Duration {
	valStr := vars.mandatory(key)

	duration, err := time.ParseDuration(valStr)
	if err != nil {
		vars.malformed = append(vars.malformed, fmt.Sprintf("mandatory %s (value=%q) is not a Duration", key, valStr))
		return 0
	}

	return duration
}

func (vars *confVars) mandatoryInt(key string) int {
	valStr := vars.mandatory(key)

	val, err := strconv.Atoi(valStr)
	if err != nil {
		vars.malformed = append(vars.malformed, fmt.Sprintf("mandatory %s (value=%q) is not a boolean", key, valStr))
		return 0
	}

	return val
}
func (vars confVars) Error() error {
	if len(vars.missing) > 0 {
		return fmt.Errorf("missing mandatory configurations: %s", strings.Join(vars.missing, ", "))
	}

	if len(vars.malformed) > 0 {
		return fmt.Errorf("malformed configurations: %s", strings.Join(vars.malformed, "; "))
	}
	return nil
}
