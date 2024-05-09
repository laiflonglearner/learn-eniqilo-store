package config

import (
	"fmt" // provides functions for formatting and printing strings.

	"github.com/backend-magang/eniqilo-store/utils/pkg" // an internal package (pkg) providing a SQL transaction service.
	"github.com/spf13/viper"                            // Viper is a popular Go library for working with configuration files and env variables.
)

type Config struct {
	// CONFIGURATION STRUCT
	AppHost    string `mapstructure:"APP_HOST"`
	AppPort    string `mapstructure:"APP_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBParams   string `mapstructure:"DB_PARAMS"`
	DBSchema   string `mapstructure:"DB_SCHEMA"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
	BCryptSalt string `mapstructure:"BCRYPT_SALT"`
	SqlTrx     *pkg.SqlWithTransactionService

	/* This struct defines the configuration for the application.
	It contains fields for various settings such as application host, port, database connection details, JWT secret, etc.
	The mapstructure tags are used to map configuration keys to struct fields when unmarshaling configuration data.
	(just in case you forget what 'marshal' and 'unmarshal' means go to common-concepts notes on discord*/
}

func Load() (conf Config) {
	viper.SetConfigFile("env")
	viper.SetConfigFile("./.env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

	return

	/* lOAD FUNCTION: Here, Load() is a function responsible for loading configuration settings.
	It first sets the paths for configuration files (env and .env).
	Then, it reads the configuration files using viper and unmarshals the configuration data into a Config struct.
	If any error occurs during reading or unmarshaling, it panics. */
}

func (cfg *Config) GetDSN() (dsn string) {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?%s&search_path=%s",
		cfg.DBUsername,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBParams,
		cfg.DBSchema,
	)

	/* GetDSN method:  a method to generate a DSN string for connecting to a PostgreSQL database
	This method, GetDSN(), is associated with the Config struct.
	It generates a Database Source Name (DSN) string used for establishing a connection to a PostgreSQL database.
	It formats the connection string using the database-related fields from the Config struct. */

}
