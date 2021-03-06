package storage

import (
	"os"
	"time"
)

// NodeInfo describes a file or folder
type NodeInfo struct {
	Path  string
	Mode  os.FileMode
	UID   uint32
	GID   uint32
	Size  int64
	ATime time.Time
	MTime time.Time
	CTime time.Time

	Metadata map[string]string
}
