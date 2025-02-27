package infrastructure

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
	"github.com/h2non/bimg"
)

// ImaginaryRepository handles image processing using bimg
type ImaginaryRepository struct{}

// NewImaginaryRepository creates a new instance of ImaginaryRepository
func NewImaginaryRepository() *ImaginaryRepository {
	return &ImaginaryRepository{}
}

// CompressImage compresses the image to specified sizes and saves them
func (r *ImaginaryRepository) CompressImage(imageData []byte, originalSize int64) (map[string]string, error) {
	paths := make(map[string]string)
	timestamp := time.Now().Format("20060102_150405") // Readable timestamp: YYYYMMDD_HHMMSS

	// Use relative path for uploads directory
	uploadDir := "../uploads" // Relative to cmd/, points to ~/Desktop/ImageCompress/uploads
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create uploads directory: %v", err)
	}

	// Determine original size compression
	var originalQuality int
	switch {
	case originalSize > 5*1024*1024: // >5MB
		originalQuality = 50 // 50% quality
	case originalSize > 2*1024*1024: // 2-5MB
		originalQuality = 70 // 70% quality (100% - 30%)
	default: // <2MB
		originalQuality = 100 // Original quality
	}

	// Save original size
	originalPath := filepath.Join(uploadDir, fmt.Sprintf("%s_original.jpg", timestamp))
	originalImg := bimg.NewImage(imageData)
	if originalQuality < 100 {
		compressed, err := originalImg.Process(bimg.Options{Quality: originalQuality})
		if err != nil {
			return nil, fmt.Errorf("failed to compress original image: %v", err)
		}
		imageData = compressed
	}
	if err := bimg.Write(originalPath, imageData); err != nil {
		return nil, fmt.Errorf("failed to save original image: %v", err)
	}
	paths["original"] = originalPath

	// Compress to 200-300KB
	size300KBPath := filepath.Join(uploadDir, fmt.Sprintf("%s_300kb.jpg", timestamp))
	if err := r.compressToTarget(imageData, size300KBPath, 200*1024, 300*1024); err != nil {
		return nil, err
	}
	paths["300kb"] = size300KBPath

	// Compress to 10-20KB
	size20KBPath := filepath.Join(uploadDir, fmt.Sprintf("%s_20kb.jpg", timestamp))
	if err := r.compressToTarget(imageData, size20KBPath, 10*1024, 20*1024); err != nil {
		return nil, err
	}
	paths["20kb"] = size20KBPath

	// Compress to 1-5KB
	size5KBPath := filepath.Join(uploadDir, fmt.Sprintf("%s_5kb.jpg", timestamp))
	if err := r.compressToTarget(imageData, size5KBPath, 1*1024, 5*1024); err != nil {
		return nil, err
	}
	paths["5kb"] = size5KBPath

	return paths, nil
}

// compressToTarget compresses an image to fit within a target size range
func (r *ImaginaryRepository) compressToTarget(imageData []byte, path string, minSize, maxSize int64) error {
	img := bimg.NewImage(imageData)
	quality := 90

	// Get image dimensions
	size, err := img.Size()
	if err != nil {
		return fmt.Errorf("failed to get image size: %v", err)
	}

	// Adjust resolution if necessary
	targetScale := 1.0
	if int64(len(imageData)) > maxSize*2 {
		targetScale = float64(maxSize) / float64(len(imageData)) * 1.5
		newWidth := int(float64(size.Width) * targetScale)
		newHeight := int(float64(size.Height) * targetScale)
		resized, err := img.Resize(newWidth, newHeight)
		if err != nil {
			return fmt.Errorf("failed to resize image for %s: %v", filepath.Base(path), err)
		}
		imageData = resized
		img = bimg.NewImage(imageData)
	}

	// Compression loop
	for i := 0; i < 20; i++ {
		compressed, err := img.Process(bimg.Options{Quality: quality})
		if err != nil {
			return fmt.Errorf("failed to compress image to %s: %v", filepath.Base(path), err)
		}
		size := int64(len(compressed))
		log.Printf("Attempt %d: Quality=%d, Size=%d bytes for %s", i+1, quality, size, filepath.Base(path))

		if size >= minSize && size <= maxSize {
			return bimg.Write(path, compressed)
		}
		if size > maxSize {
			quality -= 10
		} else {
			quality += 5
		}
		if quality < 10 {
			log.Printf("Warning: Using fallback quality 10 for %s (size=%d bytes)", filepath.Base(path), size)
			dimensions, err := img.Size()
			if err != nil {
				return fmt.Errorf("failed to get image dimensions for %s: %v", filepath.Base(path), err)
			}
			resized, err := img.Resize(dimensions.Width/2, dimensions.Height/2)
			if err != nil {
				return fmt.Errorf("failed to resize fallback for %s: %v", filepath.Base(path), err)
			}
			compressed, err = bimg.NewImage(resized).Process(bimg.Options{Quality: 10})
			if err != nil {
				return fmt.Errorf("failed to compress fallback for %s: %v", filepath.Base(path), err)
			}
			return bimg.Write(path, compressed)
		}
		if quality > 100 {
			log.Printf("Warning: Using fallback quality 100 for %s (size=%d bytes)", filepath.Base(path), size)
			return bimg.Write(path, compressed)
		}
	}

	log.Printf("Warning: Could not achieve target size %d-%d bytes for %s, using quality 50", minSize, maxSize, filepath.Base(path))
	compressed, err := img.Process(bimg.Options{Quality: 50})
	if err != nil {
		return fmt.Errorf("failed to compress fallback for %s: %v", filepath.Base(path), err)
	}
	return bimg.Write(path, compressed)
}