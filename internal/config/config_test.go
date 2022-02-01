package config

import (
	"github.com/stretchr/testify/require"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildConfigFileNotFound(t *testing.T) {
	configFile := "../../config.yaml"
	c, err := BuildConfig(configFile)
	require.Error(t, err)
	assert.Nil(t, c)
}

func TestBuildConfigEmptyFile(t *testing.T) {
	configFile := "./testdata/empty.config.yml"
	c, err := BuildConfig(configFile)
	require.Error(t, err)
	assert.Nil(t, c)
}

func TestBuildConfigSuccessDecode(t *testing.T) {
	configFile := "../../configs/config.yml"
	c, err := BuildConfig(configFile)
	require.NoError(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, "https://shrt.io", c.Host)
	assert.Equal(t, "debug", c.LogLevel)
	assert.Equal(t, "badabum", c.JWTSecret)
	assert.Equal(t, 8083, c.Port)
	assert.Equal(t, 7, c.HashMinLength)
	assert.Equal(t, "hahash", c.HashSalt)
	assert.Equal(t, "0.0.0.0", c.DBConfig.Host)
	assert.Equal(t, "postgres", c.DBConfig.User)
	assert.Equal(t, "iniT11", c.DBConfig.Password)
	assert.Equal(t, "shortener", c.DBConfig.DBName)
	assert.Equal(t, 5432, c.DBConfig.Port)
}

func TestBuildConfigSuccessEnvProcess(t *testing.T) {
	logLevel := "info"
	port := 8083
	jwtSecret := "secret"
	host := "https://test.online"
	hashSalt := "salt"
	hashMinLength := 12
	dbHost := "0.0.12.0"
	dbUser := "user"
	dbPassword := "pass"
	dbName := "test"
	dbPort := 54322
	os.Clearenv()
	os.Setenv("APP_HOST", host)
	os.Setenv("APP_LOG_LEVEL", logLevel)
	os.Setenv("APP_PORT", strconv.Itoa(port))
	os.Setenv("APP_JWT_SECRET", jwtSecret)
	os.Setenv("APP_HASH_SALT", hashSalt)
	os.Setenv("APP_HASH_MIN_LENGTH", strconv.Itoa(hashMinLength))
	os.Setenv("APP_DB_HOST", dbHost)
	os.Setenv("APP_DB_USER", dbUser)
	os.Setenv("APP_DB_PASSWORD", dbPassword)
	os.Setenv("APP_DB_NAME", dbName)
	os.Setenv("APP_DB_PORT", strconv.Itoa(dbPort))
	configFile := "../../configs/config.yml"
	c, err := BuildConfig(configFile)
	require.NoError(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, host, c.Host)
	assert.Equal(t, logLevel, c.LogLevel)
	assert.Equal(t, jwtSecret, c.JWTSecret)
	assert.Equal(t, port, c.Port)
	assert.Equal(t, hashMinLength, c.HashMinLength)
	assert.Equal(t, hashSalt, c.HashSalt)
	assert.Equal(t, dbHost, c.DBConfig.Host)
	assert.Equal(t, dbUser, c.DBConfig.User)
	assert.Equal(t, dbPassword, c.DBConfig.Password)
	assert.Equal(t, dbName, c.DBConfig.DBName)
	assert.Equal(t, dbPort, c.DBConfig.Port)
}
