package storage

import (
	"os"
	"time"
)

type HashType uint16

const (
	UndefinedHashType HashType = iota
	CRC32
	MD5
	XXHash64
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

	Hashes   map[HashType][]byte
	Metadata map[string]string
}
