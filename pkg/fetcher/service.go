package fetcher

import (
	"net"
	"net/url"
)

// Service purpose is given a URL from URL Frontier
// identify the DNS and later fetch the ip and pass
// it to the URL Renderer. Here we can also validate
// against the robot.txt to make certain parts of the
// websites are not crawled as per the robots.txt.
type Service interface {
	GetIP(host string) (net.IP, error)
	GetDNS(host string) (string, error)
}

type fetcher struct {}


func NewFetcher() Service {
	return &fetcher{}
}

// GetIP returns an IPV4 address from given host
func (f *fetcher) GetIP(rawURL string) (net.IP, error) {
	// parse as url
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	host := parsedURL.Host
	// look up IPV4 IP address
	ipAddr := *new(net.IP)
	ips, err := net.LookupIP(host)
	if err != nil {
		return nil, err
	}
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			ipAddr = ipv4.To4()
		}
	}

	return ipAddr, nil
}

// GetDNS returns the DNS server from given host
func (f *fetcher) GetDNS(host string) (string, error) {
	return "", nil
}