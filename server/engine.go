package webserver

import (
	"assignment3/server/config"
	"assignment3/server/router"
)

func Start() {
	// engine := gin.Default()
	router.CreateRouter().Run(config.PORT)
}
