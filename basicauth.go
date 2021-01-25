package basicauth

import (
	"encoding/base64"
	"fmt"
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

	fmt.Println("isNorm", r.Header)

	aHeader := r.Header.Get("Authorization")

	fmt.Println("aHeader", aHeader)

	if aHeader == "" || !strings.HasPrefix(aHeader, "Basic ") {
		return false
	}

	if aHeader != "Basic "+base64.StdEncoding.EncodeToString([]byte(d.UserName+":"+d.Password)) {
		return false
	}

	return true
}
