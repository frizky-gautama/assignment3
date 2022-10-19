package main

import (
	webserver "assignment3/server"

	"github.com/gin-gonic/gin"
)

type Status struct {
	water int
	wind  int
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	webserver.Start()
}
