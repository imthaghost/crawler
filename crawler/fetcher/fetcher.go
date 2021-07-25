package fetcher

import (
	"fmt"
	"net"
)

func Ip_address(url string) {
	iprecords, _ := net.LookupIP(url)
	fmt.Print(iprecords)

	for _, ip := range iprecords {
		fmt.Print(ip)
	}
}
