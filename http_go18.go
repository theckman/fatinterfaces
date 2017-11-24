// +build !go19

package fatinterfaces

import (
	"context"
	"io"
	"net"
)

// NetHTTPServer is an interface to describe an *http.Server.
type NetHTTPServer interface {
	io.Closer

	ListenAndServe() error
	ListenAndServeTLS(certFile, keyFile string) error
	Serve(l net.Listener) error
	SetKeepAlivesEnabled(v bool)
	Shutdown(ctx context.Context) error
}
