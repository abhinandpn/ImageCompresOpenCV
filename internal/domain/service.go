package domain

import "github.com/abhinandpn/ImageCompress/internal/infrastructure"

// ImageService defines the interface for image processing operations
type ImageService interface {
	// CompressImage compresses the image to specified sizes and returns paths
	CompressImage(imageData []byte, originalSize int64) (map[string]string, error)
}

// imageService is the concrete implementation of ImageService
type imageService struct {
	repo *infrastructure.ImaginaryRepository
}

// NewImageService creates a new instance of ImageService
func NewImageService(repo *infrastructure.ImaginaryRepository) ImageService {
	return &imageService{repo: repo}
}

// CompressImage delegates the compression to the infrastructure layer
func (s *imageService) CompressImage(imageData []byte, originalSize int64) (map[string]string, error) {
	return s.repo.CompressImage(imageData, originalSize)
}