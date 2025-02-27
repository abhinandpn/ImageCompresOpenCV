package domain

import "time"

// Image represents the image entity with its properties
type Image struct {
	OriginalPath    string
	CompressedPaths map[string]string // Key: size label, Value: file path
	UploadTime      time.Time
	Size            int64 // Size in bytes
}

// NewImage creates a new Image instance
func NewImage(size int64) *Image {
	return &Image{
		CompressedPaths: make(map[string]string),
		UploadTime:      time.Now(),
		Size:            size,
	}
}