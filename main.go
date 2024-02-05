package main

import (
	"github.com/gin-gonic/gin"
	"github.com/huin/goupnp/dcps/internetgateway2"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/gateway-ip", func(c *gin.Context) {
		clients, _, _ := internetgateway2.NewWANIPConnection1Clients()
		ipv4, _ := clients[0].GetExternalIPAddress()
		c.JSON(200, gin.H{
			"message": "Ok",
			"data":    ipv4,
		})
	})

	r.GET("/forwarded-ports", func(c *gin.Context) {
		clients, _, _ := internetgateway2.NewWANIPConnection1Clients()
		var res = make([]gin.H, 0)
		for i := 0; true; i++ {
			remoteHost, externalPort, protocol, internalPort, internalClient, enabled, portMappingDescription, leaseDuration, err := clients[0].GetGenericPortMappingEntry(uint16(i))
			if err != nil {
				break
			}
			res = append(res, gin.H{
				"remoteHost":             remoteHost,
				"externalPort":           externalPort,
				"protocol":               protocol,
				"internalPort":           internalPort,
				"internalClient":         internalClient,
				"enabled":                enabled,
				"portMappingDescription": portMappingDescription,
				"leaseDuration":          leaseDuration,
			})
		}
		c.JSON(200, gin.H{
			"message": "Ok",
			"data":    res,
		})
	})

	r.POST("/{routerPort}/{forwardedPort}", func(c *gin.Context) {
		routerPort := c.Param("routerPort")
		forwardedPort := c.Param("forwardedPort")
		c.JSON(200, gin.H{
			"message": "Port " + routerPort + " forwarded to " + forwardedPort + " successfully!",
		})
	})

	err := r.Run()
	if err != nil {
		return
	}

}
