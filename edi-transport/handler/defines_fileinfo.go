package gsftp

import (
	"os"
	"time"

	"cloud.google.com/go/storage"
)

// SyntheticFileInfo has attributes of storage object
type SyntheticFileInfo struct {
	objAttr *storage.ObjectAttrs
	prefix  string
}

// Name gets base name of the file
func (f *SyntheticFileInfo) Name() string {
	if f.objAttr.Prefix == "" {
		return f.objAttr.Name[len(f.prefix):]
	}

	return f.objAttr.Prefix[len(f.prefix) : len(f.objAttr.Prefix)-1]
}

// Size gets length in bytes for regular files; system-dependent for others
func (f *SyntheticFileInfo) Size() int64 {
	return f.objAttr.Size
}

// Mode gets file mode bits
func (f *SyntheticFileInfo) Mode() os.FileMode {
	if f.objAttr.Prefix == "" {
		return 0777
	}
	return os.ModeDir | 0777
}

// ModTime gets modification time
func (f *SyntheticFileInfo) ModTime() time.Time {
	if f.objAttr.Prefix == "" {
		return f.objAttr.Updated
	}
	return time.Now()
}

// IsDir is an abbreviation for Mode().IsDir()
func (f *SyntheticFileInfo) IsDir() bool {
	return f.objAttr.Prefix == ""
}

// Sys is an underlying data source (can return nil)
func (f *SyntheticFileInfo) Sys() interface{} {
	return nil
}
