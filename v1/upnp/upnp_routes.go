package upnp

import (
	"github.com/gin-gonic/gin"
	"github.com/huin/goupnp/dcps/internetgateway2"
	"golang.org/x/net/context"
	"networKing/utils"
)

func Register(engine *gin.RouterGroup) {
	router := engine.Group("/upnp")
	{
		router.GET("/idg", getGatewayIP)
		router.GET("/forwarded-ports", getForwardedPorts)
		router.POST("/forward", forwardPort)
	}

}

func getGatewayIP(c *gin.Context) {
	res := *utils.GetGateway()
	ip, _ := res.GetExternalIPv4Address(context.Background())
	c.JSON(200, gin.H{
		"message": "Ok",
		"data":    ip.String(),
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
	/*type PortMapping struct {
		ExternalPort   uint16 `json:"externalPort"`
		InternalPort   uint16 `json:"internalPort"`
		InternalClient string `json:"internalClient"`
	}
	var mapping PortMapping
	err := c.BindJSON(&mapping)
	println(mapping.InternalClient, mapping.InternalPort, mapping.ExternalPort)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	clients, _, _ := internetgateway2.NewWANIPConnection2Clients()
	println(clients[0].GetExternalIPAddress())
	err = clients[0].AddPortMapping("", uint16(8086), "udp", uint16(8086), "", true, "Networking", uint32(3600))
	if err != nil {
		c.JSON(500, gin.H{
			"error":   err,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Ok",
	})*/

	dev := *utils.GetGateway()
	mapping, err := dev.AddPortMapping(context.Background(), "UDP", 8086, 8086, "ok", 3600)
	if err != nil {
		c.JSON(500, gin.H{
			"error":   err,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Ok",
		"data":    mapping,
	})

}
