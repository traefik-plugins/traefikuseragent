// Package traefikuseragent a demo plugin.
package traefikuseragent

import (
	"context"
	"log"
	"net/http"

	"github.com/mssola/user_agent"
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
	ua := user_agent.New("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.97 Safari/537.11")

	name, version := ua.Browser()
	log.Printf("@@%v@@%v@@", name, version)

	mw.next.ServeHTTP(rw, req)
}
