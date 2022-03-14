package config

type DBConfig struct {
	DbHost     string `mapstructure:"DB_HOST"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbUser     string `mapstructure:"DB_USER"`
	DbName     string `mapstructure:"DB_NAME"`
}
