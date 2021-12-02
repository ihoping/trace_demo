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
	group.POST("/get-detail", func(ctx *gin.Context) {
		ctx2 := server.GetContext(ctx)

		d, err := dao.New("trace_demo", ctx2)
		article, err := d.GetArticleDetail("1")
		_, _ = d.GetArticleDetail("2")
		_, _ = d.GetArticleDetail("3")

		client, err := redis.GetClient("trace_demo")
		client.Set(ctx2, "test", "xxxxxx", time.Second*3)
		client.Set(ctx2, "name", "张三", time.Second*3)

		ctx.JSON(200, gin.H{
			"article": article,
			"err":     err,
		})

	})
}
