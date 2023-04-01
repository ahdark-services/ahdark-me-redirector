package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"

	"github.com/ahdark-services/ahdark-me-redirector/internal/env"
)

func CORS(config *env.Config) gin.HandlerFunc {
	c := cors.Config{
		AllowOrigins:     config.Cors.AllowOrigins,
		AllowMethods:     config.Cors.AllowMethods,
		AllowHeaders:     config.Cors.AllowHeaders,
		ExposeHeaders:    config.Cors.ExposeHeaders,
		AllowCredentials: config.Cors.AllowCredentials,
	}

	if lo.SomeBy(config.Cors.AllowOrigins, func(s string) bool { return s == "*" }) {
		c.AllowAllOrigins = true
		c.AllowOrigins = nil
	}

	return cors.New(c)
}
