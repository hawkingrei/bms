package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hawkingrei/bms/model"
)

func (s *Service) Upload(c *gin.Context) {
	var u model.BenchOutput
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if err := s.dao.InsertBenchmark(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.String(200, "ok")
}
