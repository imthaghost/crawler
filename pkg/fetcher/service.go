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
	GetCNAME(host string) (string, error)
	ReverseLookup(ip string) ([]string, error)
	GetNameServer(host string) ([]*net.NS, error)
	GetMX(host string) (string, error)
	GetTXT(host string) ([]string, error)
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

// GetCNAME returns the CNAME from given host
func (f *fetcher) GetCNAME(host string) (string, error) {
	cname,  err := net.LookupCNAME(host)
	if err != nil {
		return "", err
	}

	return cname, nil

}

// ReverseLookup returns the reverse record from a given ipv4 address
func (f *fetcher) ReverseLookup(addr string) ([]string, error) {
	ptr, err := net.LookupAddr(addr)
	if err !=nil {
		return nil, err
	}

	return ptr, nil
}

// GetNameServer returns nameserver records from given host
func (f *fetcher) GetNameServer(host string) ([]*net.NS, error) {
	ns, err := net.LookupNS(host)
	if err != nil {
		return nil, err
	}

	return ns, nil
}

// GetMX returns a slice of MX structs
func (f *fetcher) GetMX(host string) (string, error) {
	return "", nil
}

// GetTXT returns information about the SPF records
func (f *fetcher) GetTXT(host string) ([]string, error) {
	records, err := net.LookupTXT(host)
	if err != nil {
			return nil, err
		}

	return records, nil
}




