package fetcher

import (
	"github.com/stretchr/testify/assert"
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

// TestGetIP ...
func (suite *ServiceTestSuite) TestGetIP() {
	t := suite.T()

	test1Input := "https://golang.org"
	test1Response := "142.250.72.209"

	svc := &fetcher{}

	out1, err := svc.GetIP(test1Input)
	assert.Nil(t, err)
	assert.Equal(t, out1, test1Response)
}

// TestGetDNS ...
func (suite *ServiceTestSuite) TestGetDNS() {

}