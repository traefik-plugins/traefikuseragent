// Package traefikuseragent a demo plugin.
package traefikuseragent

import (
	"context"
	"net/http"
	"strconv"

	"github.com/mssola/user_agent"
)

const (
	// UserAgentHeader user agent header.
	UserAgentHeader = "User-Agent"
)

// Config the plugin configuration.
type Config struct{}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// TraefikUserAgent a TraefikUserAgent plugin.
type TraefikUserAgent struct {
	next http.Handler
	name string
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &TraefikUserAgent{
		next: next,
		name: name,
	}, nil
}

func (mw *TraefikUserAgent) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ua := user_agent.New(req.Header.Get(UserAgentHeader))

	req.Header.Set("X-Device-Mobile", strconv.FormatBool(ua.Mobile()))
	req.Header.Set("X-Device-Os", ua.OSInfo().Name)

	name, version := ua.Browser()
	req.Header.Set("X-Device-Browser", name)
	req.Header.Set("X-Device-Browser-Version", version)

	name, version = ua.Engine()
	req.Header.Set("X-Device-Engine", name)
	req.Header.Set("X-Device-Engine-Version", version)

	mw.next.ServeHTTP(rw, req)
}
