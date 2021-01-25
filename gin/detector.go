package gin

import (
	"net/http"

	"github.com/azubkokshe/krakend-basicauth"
	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/devopsfaith/krakend/proxy"
	krakendgin "github.com/devopsfaith/krakend/router/gin"
	"github.com/gin-gonic/gin"
)

func Register(cfg config.ServiceConfig, l logging.Logger, engine *gin.Engine) {
	credConf, err := basicauth.ParseConfig(cfg.ExtraConfig)
	if err == basicauth.ErrNoConfig {
		l.Debug("basicauth middleware: ", err.Error())
		return
	}
	if err != nil {
		l.Warning("basicauth middleware: ", err.Error())
		return
	}

	d := basicauth.New(credConf)
	engine.Use(middleware(d))
}

func New(hf krakendgin.HandlerFactory, l logging.Logger) krakendgin.HandlerFactory {
	return func(cfg *config.EndpointConfig, p proxy.Proxy) gin.HandlerFunc {
		next := hf(cfg, p)

		credCfg, err := basicauth.ParseConfig(cfg.ExtraConfig)
		if err == basicauth.ErrNoConfig {
			l.Debug("basicauth: ", err.Error())
			return next
		}
		if err != nil {
			l.Warning("basicauth: ", err.Error())
			return next
		}

		d := basicauth.New(credCfg)

		return handler(d, next)
	}
}

func middleware(f basicauth.AuthFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !f(c.Request) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func handler(f basicauth.AuthFunc, next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !f(c.Request) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		next(c)
	}
}
