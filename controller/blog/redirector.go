package blog

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/ahdark-services/ahdark-me-redirector/pkg/wpgo"
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

	post, err := ctr.WpGo.Post.Get(ctx, id, wpgo.GetPostParams{})
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to get post")
		c.String(500, err.Error())
		return
	}

	c.Redirect(301, post.Link)
}
