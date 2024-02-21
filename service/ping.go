package service

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) Ping(c *gin.Context) {
	c.String(200, "ok")
}
