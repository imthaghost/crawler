package fetcher

import (
	"net"
	"net/url"
)

func Ip_address(input string) (string, error) {
	// parse as url
	url, err := url.Parse(input)
	if err != nil {
		return "", err
	}
	host := url.Host
	iprecords, err := net.LookupIP(host)
	if err != nil {
		return "", err
	}
	ip := iprecords[1].String()

	return ip, nil
}
