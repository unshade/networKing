package general

import (
	"github.com/gin-gonic/gin"
	"net"
)

func Register(engine *gin.RouterGroup) {
	router := engine.Group("/general")
	{
		router.GET("/ping", ping)
		router.GET("/ip", ip)
	}

}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func ip(c *gin.Context) {
	ifaces, err := net.Interfaces()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	var res = make([]gin.H, 0)
	for _, iface := range ifaces {
		println("\t" + iface.Name)

		unicastAddrs, _ := iface.Addrs()
		multicastAddrs, _ := iface.MulticastAddrs()

		var unicasts = make([]string, 0)
		for _, addr := range unicastAddrs {
			unicasts = append(unicasts, addr.String())
		}
		var multicasts = make([]string, 0)
		for _, addr := range multicastAddrs {
			multicasts = append(multicasts, addr.String())
		}

		res = append(res, gin.H{
			"name":               iface.Name,
			"unicastAddresses":   unicastAddrs,
			"multicastAddresses": multicastAddrs,
		})
	}

	c.JSON(200, gin.H{
		"message": "Ok",
		"data":    res,
	})
}
