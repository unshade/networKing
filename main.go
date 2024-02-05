package main

import (
	"github.com/gin-gonic/gin"
	"networKing/v1/general"
	"networKing/v1/upnp"
)

func main() {
	engine := gin.Default()
	v1engine := engine.Group("/api/v1")
	upnp.Register(v1engine)
	general.Register(v1engine)

	err := engine.Run()
	if err != nil {
		return
	}
}
