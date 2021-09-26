package cache

import (
	"fmt"
	"time"
)

// Service caching service
type Service interface {
	Set(key string, data interface{}, expiration time.Duration) error
	Get(key string) ([]byte, error)
}


type redis struct {}

// New ...
func New() Service {
	return &redis{}
}


// Set ...
func (r *redis) Set(key string, data interface{}, expiration time.Duration) error {
	fmt.Println(key)
	fmt.Println(data)
	fmt.Println(expiration)

	return nil
}

// Get ...
func (r *redis) Get(key string) ([]byte, error) {
	return nil, nil
}