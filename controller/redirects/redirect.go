package redirects

import (
	"github.com/gin-gonic/gin"
)

func (ctr *Controller) RedirectHandler(url string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, span := tracer.Start(c.Request.Context(), "telegram-handler")
		defer span.End()
		c.Request = c.Request.WithContext(ctx)

		c.Redirect(302, url)
	}
}
