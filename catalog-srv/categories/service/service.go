package service

import "github.com/amorindev/go-microservices-app/catalog-srv/categories/port"

type service struct {
	Repo port.CategoryRepo
}

func NewSrv(repo port.CategoryRepo) *service {
	return &service{Repo: repo}
}