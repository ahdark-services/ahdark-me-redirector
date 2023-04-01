package env

import (
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
)

var (
	tracer     = otel.Tracer("internal.env")
	InstanceID = uuid.New().String()
)

type system struct {
	Name        string `default:"ahdark-me-redirector"`
	Debug       bool   `envconfig:"DEBUG" default:"false"`
	TracerDSN   string `envconfig:"TRACER_DSN" default:"http://localhost:14268/api/traces"`
	CacheDriver string `envconfig:"CACHE_DRIVER" default:"memory"`
}

type log struct {
	Driver   string `default:"console"`                                                     // console, file
	FilePath string `envconfig:"LOG_FILE_PATH" default:"/var/log/ahdark-me-redirector.log"` // only used when Driver is `file`
}

type server struct {
	Listen    string `default:":8080"`
	UnixSock  string `envconfig:"SERVER_UNIX_SOCK" default:""`
	CertFile  string `envconfig:"SERVER_CERT_FILE" default:""`
	KeyFile   string `envconfig:"SERVER_KEY_FILE" default:""`
	SSLListen string `envconfig:"SERVER_SSL_LISTEN" default:""`
}

type redis struct {
	Host     string `default:"localhost"`
	Port     int    `default:"6379"`
	Username string `default:""`
	Password string `default:""`
	DB       int    `default:"0"`
}

type cors struct {
	AllowOrigins     []string `default:"*" envconfig:"CORS_ALLOW_ORIGINS" split_words:"true"`
	AllowMethods     []string `default:"GET,HEAD,OPTIONS" envconfig:"CORS_ALLOW_METHODS" split_words:"true"`
	AllowHeaders     []string `default:"" envconfig:"CORS_ALLOW_HEADERS" split_words:"true"`
	ExposeHeaders    []string `default:"" envconfig:"CORS_EXPOSE_HEADERS" split_words:"true"`
	AllowCredentials bool     `default:"true" envconfig:"CORS_ALLOW_CREDENTIALS"`
}

type custom struct {
	RedirectConfig    string `envconfig:"CUSTOM_REDIRECT_CONFIG" default:"redirects.json"`
	WordPressEndpoint string `envconfig:"CUSTOM_WORDPRESS_ENDPOINT"`
}

type Config struct {
	System system
	Log    log
	Server server
	Redis  redis
	Cors   cors
	Custom custom
}
