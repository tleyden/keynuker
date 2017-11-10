package keynuker

import (
	"log"
	"strings"
	"testing"

	"github.com/dustin/go-humanize"
	"github.com/stretchr/testify/assert"
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

func TestCalculateBytesScanned(t *testing.T) {

	testInput := []string{
		`2017-11-06T22:27:11.029878524Z stderr: 2017/11/06 22:27:11 Scanning 3732 bytes of content for commit: ad0abf5e36350f5975ca2992c4b7510baecc45e6 url: https://api.github.com/repos/couchbase/build-team-manifests/commits/ad0abf5e36350f5975ca2992c4b7510baecc45e6`,
		`2017-11-06T22:27:11.030709919Z stderr: 2017/11/06 22:27:11 SetCheckpointIfMostRecent setting checkpoint from 6816684954 -> to 6816912337 for user: github.User{Login:\"cb-robot\"}`,
		`2017-11-06T22:27:11.030734829Z stderr: 2017/11/06 22:27:11 Fetching downstream content for event. User: cb-robot. Event id: 6817049989  Event created at: 2017-11-06 21:49:18 +0000 UTC Stored checkpoint: 2017-11-05 04:19:25 +0000 UTC Checkpoint ID: 6810349994`,
		`2017-11-06T22:27:11.030748559Z stderr: 2017/11/06 22:27:11 Scanning event. User: cb-robot. Event id: 6817049989  Event created at: 2017-11-06 21:49:18 +0000 UTC Stored checkpoint: 2017-11-05 04:19:25 +0000 UTC Checkpoint ID: 6810349994`,
		`2017-11-06T22:27:11.030758939Z stderr: 2017/11/06 22:27:11 Getting content for commit: c034d251c002d84059469d8cde31956db31001a5 url: https://api.github.com/repos/couchbase/build-team-manifests/commits/c034d251c002d84059469d8cde31956db31001a5`,
		`2017-11-06T22:27:11.095995351Z stderr: 2017/11/06 22:27:11 Scanning 1251 bytes of content for commit: c034d251c002d84059469d8cde31956db31001a5 url: https://api.github.com/repos/couchbase/build-team-manifests/commits/c034d251c002d84059469d8cde31956db31001a5`,
	}

	bytesScanned, err := CalculateBytesScanned(testInput)
	assert.NoError(t, err, "Unexpected error")
	assert.Equal(t, int(bytesScanned), (3732 + 1251))

	log.Printf("Raw content scanned: %s", humanize.Bytes(uint64(bytesScanned)))

}
