package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

// Config contains application settings
type Config struct {
	JWTSecret     string   `yaml:"jwt_secret" envconfig:"APP_JWT_SECRET"`
	LogLevel      string   `yaml:"log_level" envconfig:"APP_LOG_LEVEL"`
	Port          int      `yaml:"port" envconfig:"PORT"`
	HashSalt      string   `yaml:"hash_salt" envconfig:"APP_HASH_SALT"`
	HashMinLength int      `yaml:"hash_min_length" envconfig:"APP_HASH_MIN_LENGTH"`
	Host          string   `yaml:"host" envconfig:"APP_HOST"`
	DBConfig      DBConfig `yaml:"db_config"`
}

// DBConfig contains database connection settings
type DBConfig struct {
	Host     string `yaml:"host" envconfig:"APP_DB_HOST"`
	User     string `yaml:"user" envconfig:"APP_DB_USER"`
	Password string `yaml:"password" envconfig:"APP_DB_PASSWORD"`
	DBName   string `yaml:"db_name" envconfig:"APP_DB_NAME"`
	Port     int    `yaml:"port" envconfig:"APP_DB_PORT"`
}

// BuildConfig creates a configuration structure using a config file
// in .yml format and also overrides base settings with environment variables
func BuildConfig(configPath string) (*Config, error) {
	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// create config structure
	var c Config

	// Start YAML decoding from file
	if err := d.Decode(&c); err != nil {
		return nil, err
	}

	// If any parameter is passed through environment variables,
	// then the passed value must override the value from the configuration file.
	err = envconfig.Process("", &c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
