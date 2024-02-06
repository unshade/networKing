package utils

import (
	"github.com/syncthing/syncthing/lib/nat"
	"github.com/syncthing/syncthing/lib/upnp"
	"golang.org/x/net/context"
)

func GetGateway() *nat.Device {
	background := context.Background()
	discover := upnp.Discover(background, 1000000000, 800000000)
	var res *nat.Device
	for _, device := range discover {
		_, err := device.GetExternalIPv4Address(background)
		if err != nil {
			println(err.Error())
		} else {
			res = &device
			break
		}
	}
	return res
}
