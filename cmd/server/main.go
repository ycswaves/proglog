package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ycswaves/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer()
	router := gin.Default()
	router.POST("/produce", srv.HandleProduce)
	router.GET("/consume", srv.HandleConsume)
	router.Run()
}
