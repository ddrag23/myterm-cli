package entity

import (
	"io/fs"
	"time"
)

type SSHCredentials struct {
	Label    string
	Address  string
	Port     int
	Username string
	Password string
}

// IsDir implements fs.FileInfo.
func (SSHCredentials) IsDir() bool {
	panic("unimplemented")
}

// ModTime implements fs.FileInfo.
func (SSHCredentials) ModTime() time.Time {
	panic("unimplemented")
}

// Mode implements fs.FileInfo.
func (SSHCredentials) Mode() fs.FileMode {
	panic("unimplemented")
}

// Name implements fs.FileInfo.
func (SSHCredentials) Name() string {
	panic("unimplemented")
}

// Size implements fs.FileInfo.
func (SSHCredentials) Size() int64 {
	panic("unimplemented")
}

// Sys implements fs.FileInfo.
func (SSHCredentials) Sys() any {
	panic("unimplemented")
}
