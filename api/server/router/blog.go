package router

import (
	"github.com/gin-gonic/gin"
	"trace_demo/pkg/dao"
	"trace_demo/pkg/server"
)

type Blog struct{}

func (blog Blog) Route(group *gin.RouterGroup) {
	group.GET("/get-detail", func(ctx *gin.Context) {
		ctx2 := server.GetContext(ctx)

		d := dao.New(ctx2, "trace_demo")
		article, err := d.GetArticleDetail("1")
		_, _ = d.GetArticleDetail("2")

		ctx.JSON(200, gin.H{
			"article": article,
			"err":     err,
		})

	})
}
