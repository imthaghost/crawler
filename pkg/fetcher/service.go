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
	GetIP(host string) (string, error)
	GetDNS(host string) (string, error)
}

type fetcher struct {}


func NewFetcher() Service {
	return &fetcher{}
}

// GetIP returns an IPV4 address from given host
func (f *fetcher) GetIP(host string) (string, error) {
	// parse as url
	u, err := url.Parse(host)
	if err != nil {
		return "", err
	}
	h := u.Host
	records, err := net.LookupIP(h)
	if err != nil {
		return "", err
	}
	// TODO: better implementation
	ip := records[1].String()

	return ip, nil
}

// GetDNS returns the DNS server from given host
func (f *fetcher) GetDNS(host string) (string, error) {
	return "", nil
}