package config

var Env = struct {
	// Authentication
	OAuthClientID     string `env:"OAUTH_CLIENT_ID"`
	OAuthClientSecret string `env:"OAUTH_CLIENT_SECRET"`
	// MiddlewareVerboseLoggerEnabled - to enable the negroni-logrus logger for all the endpoints useful for debugging
	MiddlewareVerboseLoggerEnabled bool `env:"MIDDLEWARE_VERBOSE_LOGGER_ENABLED" envDefault:"true"`
	// MiddlewareVerboseLoggerExcludeURLs - to exclude urls from the verbose logger via comma separated list
	MiddlewareVerboseLoggerExcludeURLs []string `env:"MIDDLEWARE_VERBOSE_LOGGER_EXCLUDE_URLS" envDefault:"/esv/health" envSeparator:","`
}{}
