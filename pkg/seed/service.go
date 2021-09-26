package seed


type Service interface {
	RemoveAbsolute(string) (string, error)
}

type seed struct {}


// RemoveAbsolute removes an absolute URL from the shared URL
func (s *seed) RemoveAbsolute(rawURL string) (string, error) {
	return "", nil
}