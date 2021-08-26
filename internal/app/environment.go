package app

// Environment holds the environment variables used by the service
type Environment struct {
	LogLevel        string `env:"LOG_LEVEL" envDefault:"debug"`
	Env             string `env:"ENV" envDefault:"dev"`
	MySQLHost       string `env:"MYSQL_HOST,required"`
	MySQLUser       string `env:"MYSQL_USER,required"`
	MySQLPassword   string `env:"MYSQL_PASSWORD,required"`
	Timezone        string `env:"TZ" envDefault:"Europe/London"`
	GuardianBaseURL string `env:"GUARDIAN_BASE_URL,required"`
	GuardianKey     string `env:"GUARDIAN_KEY,required"`
	DiscordWebhook  string `env:"DISCORD_WEBHOOK"`
	AllowOrigin     string `env:"ALLOW_ORIGIN"`
	RawDataPath     string `env:"RAW_DATA_PATH,required"`
}
