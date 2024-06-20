package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
	"time"
)

var once sync.Once

type Config struct {
	Host         string        `env:"HOST" env-required:"true"`
	Port         uint16        `env:"PORT" env-required:"true"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT" env-required:"true"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" env-required:"true"`
	IdleTimeout  time.Duration `env:"IDLE_TIMEOUT" env-required:"true"`

	AccessTokenTTL  time.Duration `env:"ACCESS_TOKEN_TTL" env-required:"true"`
	RefreshTokenTTL time.Duration `env:"REFRESH_TOKEN_TTL" env-required:"true"`
	JWTSignedKey    []byte        `env:"JWT_SIGNED_KEY" env-required:"true"`

	DBConnectionString string `env:"DB_CONNECTION_STRING" env-required:"true"`

	GeminiAPIKey string `env:"GEMINI_API_KEY" env-required:"true"`

	SmtpHost     string `env:"SMTP_HOST" env-required:"true"`
	SmtpPort     int    `env:"SMTP_PORT" env-required:"true"`
	SenderEmail  string `env:"SENDER_EMAIL" env-required:"true"`
	SMTPPassword string `env:"SMTP_APP_PASSWORD" env-required:"true"`
}

func New(filePath string) (*Config, error) {
	var cfg Config
	var err error

	once.Do(func() {
		err = cleanenv.ReadConfig(filePath, &cfg)
	})

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
