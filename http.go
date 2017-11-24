package fatinterfaces

import (
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

// NetHTTPClient is an interface to describe an *http.Client.
type NetHTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
	Get(url string) (*http.Response, error)
	Head(url string) (*http.Response, error)
	Post(url string, contentType string, body io.Reader) (*http.Response, error)
	PostForm(url string, data url.Values) (*http.Response, error)
}

// NetHTTPHeader is an interface to describe an http.Header.
type NetHTTPHeader interface {
	Add(key, value string)
	Del(key string)
	Get(key string) string
	Set(key, value string)
	Write(w io.Writer) error
	WriteSubset(w io.Writer, exclude map[string]bool) error
}

// NetHTTPRequest is an interface to describe an *http.Request.
type NetHTTPRequest interface {
	AddCookie(c *http.Cookie)
	BasicAuth() (username, password string, ok bool)
	Context() context.Context
	Cookie(name string) (*http.Cookie, error)
	Cookies() []*http.Cookie
	FormFile(key string) (multipart.File, *multipart.FileHeader, error)
	FormValue(key string) string
	MultipartReader() (*multipart.Reader, error)
	ParseForm() error
	ParseMultipartForm(maxMemory int64) error
	PostFormValue(key string) string
	ProtoAtLeast(major, minor int) bool
	Referer() string
	SetBasicAuth(username, password string)
	UserAgent() string
	WithContext(ctx context.Context) *http.Request
	Write(w io.Writer) error
	WriteProxy(w io.Writer) error
}

// NetHTTPResponse is an interface to describe an *http.Response.
type NetHTTPResponse interface {
	Cookies() []*http.Cookie
	Location() (*url.URL, error)
	ProtoAtLeast(major, minor int) bool
	Write(w io.Writer) error
}

// NetHTTPServeMux is an interface to describe an *http.ServeMux.
type NetHTTPServeMux interface {
	Handle(pattern string, handler http.Handler)
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	Handler(r *http.Request) (h http.Handler, pattern string)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
