package config

import (
	"errors"
	"log/slog"
	"os"
	"strconv"
)

var (
	ErrInvalidDbUrl = errors.New("db url is invalid")
	ErrInvalidPort  = errors.New("port number is invalid")
	ErrLogLevel     = errors.New("log level not set or invalid")
	ErrInvalidEnv   = errors.New("env not set or invliad")
	ErrRegistryURI  = errors.New("registry uri is invalid")
	ErrHostName     = errors.New("couldn't get host name from docker")
)

type Config struct {
	Port        int
	DbUrl       string
	Env         string
	LogLevel    slog.Level
	RegistryURI string
	HostName    string
	//
	SMTPhost     string
	SMTPport     int
	SMTPusername string
	SMTPpassword string
}

type Option func(*Config) error

var Environment = map[string]string{
	"dev":  "dev",
	"prod": "prod",
}
var LogLevels = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

func WithEnv() Option { // dev or prod
	return func(c *Config) error {
		env := os.Getenv("ENV")
		if env == "" {
			return ErrInvalidEnv
		}
		evalue, ok := Environment[env]
		if !ok {
			return ErrInvalidEnv
		}
		c.Env = evalue
		return nil
	}
}

func WithDbUrl() Option {
	return func(c *Config) error {
		dbUrl := os.Getenv("DB_URL")
		if dbUrl == "" {
			return ErrInvalidDbUrl
		}
		c.DbUrl = dbUrl
		return nil
	}
}
func WithSMTPHost() Option {
	return func(c *Config) error {
		smtpHost := os.Getenv("SMTP_HOST")
		if smtpHost == "" {
			return errors.New("SMTP host is required")
		}
		c.SMTPhost = smtpHost
		return nil
	}
}

func WithSMTPPort() Option {
	return func(c *Config) error {
		smtpPortStr := os.Getenv("SMTP_PORT")
		if smtpPortStr == "" {
			return errors.New("SMTP port is required")
		}
		smtpPort, err := strconv.Atoi(smtpPortStr)
		if err != nil {
			return errors.New("invalid SMTP port")
		}
		c.SMTPport = smtpPort
		return nil
	}
}

func WithSMTPUsername() Option {
	return func(c *Config) error {
		smtpUsername := os.Getenv("SMTP_USERNAME")
		if smtpUsername == "" {
			return errors.New("SMTP username is required")
		}
		c.SMTPusername = smtpUsername
		return nil
	}
}

func WithSMTPPassword() Option {
	return func(c *Config) error {
		smtpPassword := os.Getenv("SMTP_PASSWORD")
		if smtpPassword == "" {
			return errors.New("SMTP password is required")
		}
		c.SMTPpassword = smtpPassword
		return nil
	}
}
func WithPort() Option {
	return func(c *Config) error {
		portStr := os.Getenv("PORT")
		if portStr == "" {
			return ErrInvalidPort
		}
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return ErrInvalidPort
		}
		c.Port = port
		return nil
	}
}

func WithLogLevel() Option {
	return func(c *Config) error {
		logLevel := os.Getenv("LOG_LEVEL")
		if logLevel == "" {
			return ErrLogLevel
		}

		lvl, ok := LogLevels[logLevel]
		if !ok {
			return ErrLogLevel
		}
		c.LogLevel = lvl
		return nil
	}
}

func WithRegistryURI() Option {
	return func(c *Config) error {
		registryURI := os.Getenv("REGISTRY_URI")
		if registryURI == "" {
			return ErrRegistryURI
		}
		c.RegistryURI = registryURI
		return nil
	}
}
func NewConfig(options ...Option) (*Config, error) {
	config := &Config{}
	for _, opt := range options {
		if err := opt(config); err != nil {
			return nil, err
		}
	}
	return config, nil
}
