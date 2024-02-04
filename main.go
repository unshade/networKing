package main

import "github.com/gin-gonic/gin"
import "github.com/syncthing/syncthing/lib/upnp"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()

	r.POST("/{routerPort}/{forwardedPort}", func(c *gin.Context) {
		routerPort := c.Param("routerPort")
		forwardedPort := c.Param("forwardedPort")
		c.JSON(200, gin.H{
			"message": "Port " + routerPort + " forwarded to " + forwardedPort + " successfully!",
		})
	})
}
