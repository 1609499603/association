package utils

import (
	"association/global"
	"net"
)

func GetLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		global.ASS_LOG.Error("ip acquisition failed,error:" + err.Error())
		return ""
	}
	for _, address := range addrs {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return ""
}
