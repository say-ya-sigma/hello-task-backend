package controller

import "github.com/gin-gonic/gin"

type Hello interface {
	Get(ctx *gin.Context) (*ResponseGetHello, error)
}

type ResponseGetHello struct {
	Say string
}

type hello struct {
}

func CreateHello() Hello {
	return &hello{}
}

func (h hello) Get(ctx *gin.Context) (*ResponseGetHello, error) {
	return &ResponseGetHello{
		Say: "hello",
	}, nil
}

