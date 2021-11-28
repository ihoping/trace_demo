package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"trace_demo/pkg/dao"
	"trace_demo/pkg/dao/blog_dao"
)

type Blog struct{}

func (blog Blog) Route(group *gin.RouterGroup) {
	group.GET("/get-detail", func(ctx *gin.Context) {
		ctx2 := context.WithValue(ctx, "request_id", "lalalal")
		db := dao.GetDB(ctx2, "trace_demo")
		rs := &blog_dao.Blog{}
		db.Model(rs).Where("id", "1").First(rs)

		ctx.JSON(200, gin.H{
			"rs": rs.Title,
		})

	})
}
