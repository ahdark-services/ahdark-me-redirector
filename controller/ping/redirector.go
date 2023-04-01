package ping

import (
	"github.com/gin-gonic/gin"
)

func (ctr *Controller) PingHandler(c *gin.Context) {
	ctx, span := tracer.Start(c.Request.Context(), "redirect-post-handler")
	defer span.End()
	c.Request = c.Request.WithContext(ctx)

	c.String(200, "pong")
}
