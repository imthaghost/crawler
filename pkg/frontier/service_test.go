package frontier


import (
	"github.com/stretchr/testify/suite"
	"testing"
)
// ServiceTestSuite instantiate testing suite
type ServiceTestSuite struct {
	suite.Suite
}

// TestServiceTestSuite ...
func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}



