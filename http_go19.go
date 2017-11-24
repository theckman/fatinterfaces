// +build go19

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
	RegisterOnShutdown(f func())
	Serve(l net.Listener) error
	ServeTLS(l net.Listener, certFile, keyFile string) error
	SetKeepAlivesEnabled(v bool)
	Shutdown(ctx context.Context) error
}
