package server

import (
	"github.com/gin-gonic/gin"
	"trace_demo/api/server/router"
	"trace_demo/configs"
	"trace_demo/pkg/dao"
	"trace_demo/pkg/redis"
)

type Server struct {
}

func (s *Server) Startup(addr string) error {
	r := gin.Default()

	router.Blog{}.Route(r.Group("/blog"))

	err := dao.Startup([]dao.Config{
		{
			Name:     configs.MySqlTraceDemo.Name,
			Host:     configs.MySqlTraceDemo.Host,
			Port:     configs.MySqlTraceDemo.Port,
			User:     configs.MySqlTraceDemo.User,
			Password: configs.MySqlTraceDemo.Password,
			Database: configs.MySqlTraceDemo.Database,
		},
	})
	if err != nil {
		return err
	}

	redis.Startup([]redis.Config{
		{
			Name:     configs.RedisTraceDemo.Name,
			Host:     configs.RedisTraceDemo.Host,
			Port:     configs.RedisTraceDemo.Port,
			Password: configs.RedisTraceDemo.Password,
			DB:       configs.RedisTraceDemo.DB,
		},
	})

	return r.Run(addr)
}

func (s *Server) Shutdown() error {
	err := dao.Shutdown()
	if err != nil {
		return err
	}
	return redis.Shutdown()
}
