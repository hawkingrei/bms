package service

import (
	"github.com/gin-gonic/gin"
	"github.com/hawkingrei/bms/config"
	"github.com/hawkingrei/bms/dao"
)

type Service struct {
	dao  *dao.DAO
	http *gin.Engine
}

func NewService(cfg *config.Configure) (*Service, error) {
	dao, err := dao.NewDAO(cfg)
	if err != nil {
		return nil, err
	}
	srv := &Service{
		dao:  dao,
		http: gin.Default(),
	}
	srv.init()
	return srv, nil
}

func (s *Service) init() {
	s.http.GET("/ping", s.Ping)
}

func (s *Service) Run() {
	s.http.Run(":8080")
}
