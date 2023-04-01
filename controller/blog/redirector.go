package blog

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (ctr *Controller) RedirectPostHandler(c *gin.Context) {
	ctx, span := tracer.Start(c.Request.Context(), "redirect-post-handler")
	defer span.End()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to parse id")
		c.String(400, err.Error())
		return
	}

	url, err := ctr.BlogSvc.GetPostUrl(ctx, id)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to get post url")
		c.String(500, err.Error())
		return
	}

	c.Redirect(301, url)
}
