package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

// App config struct
type Config struct {
	Server    ServerConfig
	Auth      AuthConfig
	Postgres  PostgresConfig
	Redis     RedisConfig
	MongoDB   MongoDBConfig
	Cookie    CookieConfig
	Store     StoreConfig
	Session   SessionConfig
	Mailer    MailerConfig
	Email     EmailConfig
	Metrics   MetricsConfig
	Logger    LoggerConfig
	AWS       AWSConfig
	Jaeger    JaegerConfig
	Portfolio PortfolioConfig
	TCS       TCSConfig
	CBR       CBRConfig
}

// Server config struct
type ServerConfig struct {
	AppVersion        string
	Port              string
	PprofPort         string
	Mode              string
	FrontendUrl       string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	SSL               bool
	CtxDefaultTimeout time.Duration
	CSRF              bool
	CsrfSalt          string
	Debug             bool
}

func (s ServerConfig) IsDevelopment() bool {
	return s.Mode == "Development"
}

// Auth config
type AuthConfig struct {
	PrefixAccessToken    string
	AccessSecretKey      string
	AccessTokenExpMinute time.Duration
	MaxRefreshSession    int
	NameRefreshToken     string
	RefreshSecretKey     string
}

// Logger config
type LoggerConfig struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// Postgresql config
type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  bool
	PgDriver           string
}

// Redis config
type RedisConfig struct {
	RedisAddr      string
	RedisPassword  string
	RedisDB        string
	RedisDefaultdb string
	MinIdleConns   int
	PoolSize       int
	PoolTimeout    int
	Password       string
	DB             int
}

// MongoDB config
type MongoDBConfig struct {
	MongoURI string
}

// Cookie config
type CookieConfig struct {
	MaxAge   int
	Secure   bool
	HTTPOnly bool
}

// Session config
type SessionConfig struct {
	Prefix string
	Name   string
	Expire int
}

// Mailer config
type MailerConfig struct {
	Host       string
	Port       int
	User       string
	Password   string
	Encryption string
	FromEmail  string
	Mechanism  string
}

type EmailConfig struct {
	BaseLayout       string
	ConfirmedPartial string
}

// Metrics config
type MetricsConfig struct {
	URL         string
	ServiceName string
}

// Store config
type StoreConfig struct {
	ImagesFolder string
}

// AWS S3
type AWSConfig struct {
	Endpoint       string
	MinioAccessKey string
	MinioSecretKey string
	UseSSL         bool
	MinioEndpoint  string
}

// Jaeger config
type JaegerConfig struct {
	Host        string
	ServiceName string
	LogSpans    bool
}

// Portfolio config
type PortfolioConfig struct {
	TitleDefault    string
	CurrencyDefault string
}

// TCS config
type TCSConfig struct {
	Token string
}

// CBR config
type CBRConfig struct {
	SourceCurrency string
}

// Create config
func NewConfig(path string) (*Config, error) {
	cfgFile, err := LoadConfig(path)

	if err != nil {
		return nil, err
	}

	return ParseConfig(cfgFile)
}

// Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath("./config")
	v.SetConfigName(filename)
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("Config file not found")
		}
		return nil, err
	}
	return v, nil
}

// Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var config Config

	err := v.Unmarshal(&config)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
		return nil, err
	}

	return &config, nil
}
