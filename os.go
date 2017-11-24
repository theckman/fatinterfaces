package fatinterfaces

import (
	"io"
	"os"
)

// OsFile is the interface describing an *os.File.
type OsFile interface {
	io.Closer
	io.Seeker
	io.Reader
	io.ReaderAt
	io.Writer
	io.WriterAt

	Chdir() error
	Chmod(mode os.FileMode) error
	Chown(uid, gid int) error
	Fd() uintptr
	Name() string
	Readdir(n int) ([]os.FileInfo, error)
	Readdirnames(n int) (names []string, err error)
	Stat() (os.FileInfo, error)
	Sync() error
	Truncate(size int64) error
	WriteString(s string) (n int, err error)
}
