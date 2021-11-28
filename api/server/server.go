package server

import (
	"github.com/gin-gonic/gin"
	"trace_demo/api/server/router"
	"trace_demo/configs"
	"trace_demo/pkg/dao"
)

func Run(addr string) error {
	r := gin.Default()

	router.Blog{}.Route(r.Group("/blog"))

	dao.Startup(dao.Config{DSNList: []configs.MysqlDSN{
		configs.MySqlTraceDemo,
	}})
	return r.Run(addr)
}
