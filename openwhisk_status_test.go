package keynuker

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

func TestCreateApiHostBaseUrl(t *testing.T) {

	hostNoSchemeOrPort := "http://foo.com"

	hostSchemeAndPort := "https://foo.com:443"

	baseUrl, err := CreateApiHostBaseUrl(hostNoSchemeOrPort)
	assert.NoError(t, err, "Unexpected err")
	assert.True(t, baseUrl.Scheme == "http")
	assert.True(t, baseUrl.Port() == "")
	assert.True(t, strings.Contains(baseUrl.Path, "api"))

	baseUrl, err = CreateApiHostBaseUrl(hostSchemeAndPort)
	assert.NoError(t, err, "Unexpected err")
	assert.True(t, baseUrl.Scheme == "https")
	assert.True(t, baseUrl.Port() == "443")
	assert.True(t, strings.Contains(baseUrl.Path, "api"))


}