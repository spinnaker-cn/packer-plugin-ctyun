package core

import "time"

type Config struct {
	Scheme   string
	Endpoint string
	Timeout  time.Duration
}

// NewConfig returns a pointer of Config
//
// scheme only accepts http or https
//
// endpoint is the host to access, the connection could not be created if it's error
func NewConfig() *Config {
	return &Config{SchemeHttps, "ctecs-global.ctapi.ctyun.cn", 10 * time.Second}
}

func (c *Config) SetScheme(scheme string) {
	c.Scheme = scheme
}

func (c *Config) SetEndpoint(endpoint string) {
	c.Endpoint = endpoint
}

func (c *Config) SetTimeout(timeout time.Duration) {
	c.Timeout = timeout
}