package pkging

import (
	"net/http"
	"os"

	"github.com/markbates/pkger/here"
)

type File interface {
	Close() error
	Abs() (string, error)
	Info() here.Info
	Name() string
	Open(name string) (http.File, error)
	Path() here.Path
	Read(p []byte) (int, error)
	Readdir(count int) ([]os.FileInfo, error)
	Seek(offset int64, whence int) (int64, error)
	Stat() (os.FileInfo, error)
	Write(b []byte) (int, error)
}