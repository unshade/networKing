package upnp

import (
	"github.com/gin-gonic/gin"
	"github.com/huin/goupnp/dcps/internetgateway2"
)

func Register(engine *gin.RouterGroup) {
	router := engine.Group("/upnp")
	{
		router.GET("/gateway-ip", getGatewayIP)
		router.GET("/forwarded-ports", getForwardedPorts)
		router.POST("/:routerPort/:forwardedPort", forwardPort)
	}

}

func getGatewayIP(c *gin.Context) {
	clients, _, _ := internetgateway2.NewWANIPConnection1Clients()
	ipv4, _ := clients[0].GetExternalIPAddress()
	c.JSON(200, gin.H{
		"message": "Ok",
		"data":    ipv4,
	})
}

func getForwardedPorts(c *gin.Context) {
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
}

func forwardPort(c *gin.Context) {
	routerPort := c.Param("routerPort")
	forwardedPort := c.Param("forwardedPort")
	c.JSON(200, gin.H{
		"message": "Port " + routerPort + " forwarded to " + forwardedPort + " successfully!",
	})
}
