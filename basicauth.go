package basicauth

import (
	"encoding/base64"
	"net/http"
	"strings"
)

type Config struct {
	UserName string
	Password string
}

type AuthFunc func(r *http.Request) bool

func New(cfg Config) AuthFunc {
	return cfg.IsNorm
}

func (d *Config) IsNorm(r *http.Request) bool {
	aHeader := r.Header.Get("Authorization")
	if aHeader == "" || strings.HasPrefix(aHeader, "Basic ") {
		return false
	}

	if aHeader != "Basic " + base64.StdEncoding.EncodeToString([]byte(d.UserName + ":" + d.Password)) {
		return false
	}

	return true
}