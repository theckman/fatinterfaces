package fatinterfaces

import (
	"net/http"
	"testing"
)

func TestNetHTTPClient(t *testing.T) {
	var _ NetHTTPClient = (*http.Client)(nil)
}

func TestNetHTTPHeader(t *testing.T) {
	var _ NetHTTPHeader = (http.Header)(nil)
}

func TestNetHTTPRequest(t *testing.T) {
	var _ NetHTTPRequest = (*http.Request)(nil)
}

func TestNetHTTPResponse(t *testing.T) {
	var _ NetHTTPResponse = (*http.Response)(nil)
}

func TestNetHTTPServeMux(t *testing.T) {
	var _ NetHTTPServeMux = (*http.ServeMux)(nil)
}

func TestNetHTTPServer(t *testing.T) {
	var _ NetHTTPServer = (*http.Server)(nil)
}
