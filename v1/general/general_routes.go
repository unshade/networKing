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

	for _, iface := range ifaces {
		println("\t" + iface.Name)
		unicastAddrs, _ := iface.Addrs()
		multicastAddrs, _ := iface.MulticastAddrs()

		for _, addr := range unicastAddrs {
			println("unicast " + addr.String())
		}
		for _, addr := range multicastAddrs {
			println("Multicast " + addr.String())
		}
	}
}
