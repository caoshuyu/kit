package nettools

import (
	"encoding/binary"
	"net"
	"strconv"
	"strings"
)

//get local host ip v4
func GetLocalIp() (ip string,err error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return
	}
	for _, netInt := range netInterfaces {
		if (netInt.Flags & net.FlagUp) != 0 {
			addrs, er := netInt.Addrs()
			if nil != er {
				err = er
				return
			}
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						ip = ipnet.IP.String()
						return
					}
				}
			}
		}
	}
	return
}


//ip v4 to int64
func Ipv4ToInt64(ip string) (ipNum int64) {
	ipAddr, err := net.ResolveIPAddr("ip", ip)
	if err != nil {
		return 0
	}
	return int64(binary.BigEndian.Uint32(ipAddr.IP.To4()))
}

//int to ip v4
func Int64ToIpv4(ipNum int64) (ip string) {
	var bytes [4]byte
	bytes[0] = byte((ipNum >> 24) & 0xFF)
	bytes[1] = byte((ipNum >> 16) & 0xFF)
	bytes[2] = byte((ipNum >> 8) & 0xFF)
	bytes[3] = byte(ipNum & 0xFF)

	var ipBitArr []string
	for i := 0; i < 4; i++ {
		ipBitArr = append(ipBitArr, strconv.FormatInt(int64(bytes[i]), 10))
	}
	return strings.Join(ipBitArr, ".")
}
