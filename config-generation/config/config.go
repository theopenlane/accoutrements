package config

import (
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/mcuadros/go-defaults"
	"github.com/theopenlane/beacon/otelx"
	"github.com/theopenlane/core/pkg/middleware/ratelimit"
)

var (
	DefaultConfigFilePath = "./config/.config.yaml"
)

// Config contains the configuration for the openlane server
type Config struct {
	// RefreshInterval determines how often to reload the config
	RefreshInterval time.Duration `json:"refresh_interval" koanf:"refresh_interval" default:"10m"`
	// Server contains the echo server settings
	Server Server `json:"server" koanf:"server"`
	// Tracer contains the tracing config for opentelemetry
	Tracer otelx.Config `json:"tracer" koanf:"tracer"`
	// Ratelimit contains the configuration for the rate limiter
	Ratelimit ratelimit.Config `json:"ratelimit" koanf:"ratelimit"`
}

// Server settings for the echo server
type Server struct {
	// Debug enables debug mode for the server
	Debug bool `json:"debug" koanf:"debug" default:"false"`
	// Dev enables echo's dev mode options
	Dev bool `json:"dev" koanf:"dev" default:"false"`
	// Listen sets the listen address to serve the echo server on
	Listen string `json:"listen" koanf:"listen" jsonschema:"required" default:":17610"`
	// ShutdownGracePeriod sets the grace period for in flight requests before shutting down
	ShutdownGracePeriod time.Duration `json:"shutdown_grace_period" koanf:"shutdown_grace_period" default:"10s"`
	// ReadTimeout sets the maximum duration for reading the entire request including the body
	ReadTimeout time.Duration `json:"read_timeout" koanf:"read_timeout" default:"15s"`
	// WriteTimeout sets the maximum duration before timing out writes of the response
	WriteTimeout time.Duration `json:"write_timeout" koanf:"write_timeout" default:"15s"`
	// IdleTimeout sets the maximum amount of time to wait for the next request when keep-alives are enabled
	IdleTimeout time.Duration `json:"idle_timeout" koanf:"idle_timeout" default:"30s"`
	// ReadHeaderTimeout sets the amount of time allowed to read request headers
	ReadHeaderTimeout time.Duration `json:"read_header_timeout" koanf:"read_header_timeout" default:"2s"`
	// TLS contains the tls configuration settings
	TLS TLS `json:"tls" koanf:"tls"`
	// CORS contains settings to allow cross origin settings and insecure cookies
	CORS CORS `json:"cors" koanf:"cors"`
	// Openlane contains the token for the openlane server
	Openlane Openlane `json:"openlane" koanf:"openlane"`
}

// CORS settings for the server to allow cross origin requests
type CORS struct {
	// AllowOrigins is a list of allowed origin to indicate whether the response can be shared with
	// requesting code from the given origin
	AllowOrigins []string `json:"allow_origins" koanf:"allow_origins"`
	// CookieInsecure allows CSRF cookie to be sent to servers that the browser considers
	// unsecured. Useful for cases where the connection is secured via VPN rather than
	// HTTPS directly.
	CookieInsecure bool `json:"cookie_insecure" koanf:"cookie_insecure"`
}

// Openlane settings for the server to authenticate with the openlane server
type Openlane struct {
	// Token is the token used to authenticate with the openlane server
	Token string `json:"token" koanf:"token"`
}

// TLS settings for the server for secure connections
type TLS struct {
	// Config contains the tls.Config settings
	Config *tls.Config `json:"config" koanf:"config" jsonschema:"-"`
	// Enabled turns on TLS settings for the server
	Enabled bool `json:"enabled" koanf:"enabled" default:"false"`
	// CertFile location for the TLS server
	CertFile string `json:"cert_file" koanf:"cert_file" default:"server.crt"`
	// CertKey file location for the TLS server
	CertKey string `json:"cert_key" koanf:"cert_key" default:"server.key"`
	// AutoCert generates the cert with letsencrypt, this does not work on localhost
	AutoCert bool `json:"auto_cert" koanf:"auto_cert" default:"false"`
}

// Load is responsible for loading the configuration from a YAML file and environment variables.
// If the `cfgFile` is empty or nil, it sets the default configuration file path.
// Config settings are taken from default values, then from the config file, and finally from environment
// the later overwriting the former.
func Load(cfgFile *string) (*Config, error) {
	k := koanf.New(".")

	if cfgFile == nil || *cfgFile == "" {
		*cfgFile = DefaultConfigFilePath
	}

	// load defaults
	conf := &Config{}
	defaults.SetDefaults(conf)

	// parse yaml config
	if err := k.Load(file.Provider(*cfgFile), yaml.Parser()); err != nil {
		fmt.Println("failed to load config file", err)
	} else {
		// unmarshal the config
		if err := k.Unmarshal("", &conf); err != nil {
			panic(err)
		}
	}

	// load env vars
	if err := k.Load(env.Provider("OPENLANECLOUD_", ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(
			strings.TrimPrefix(s, "OPENLANECLOUD_")), "_", ".")
	}), nil); err != nil {
		panic(err)
	}

	// unmarshal the env vars
	if err := k.Unmarshal("", &conf); err != nil {
		panic(err)
	}

	return conf, nil
}
