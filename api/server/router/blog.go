package router

import (
	"github.com/gin-gonic/gin"
	"time"
	"trace_demo/pkg/dao"
	"trace_demo/pkg/redis"
	"trace_demo/pkg/server"
)

type Blog struct{}

func (blog Blog) Route(group *gin.RouterGroup) {
	group.GET("/get-detail", func(ctx *gin.Context) {
		ctx2 := server.GetContext(ctx)

		d, err := dao.New("trace_demo", ctx2)
		article, err := d.GetArticleDetail("1")
		_, _ = d.GetArticleDetail("2")

		redis, err := redis.GetClient("trace_demo", ctx2)
		redis.Set(ctx2, "test", "name", time.Second*3)

		ctx.JSON(200, gin.H{
			"article": article,
			"err":     err,
		})

	})
}
