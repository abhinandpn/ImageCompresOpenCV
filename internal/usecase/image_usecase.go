package usecase

import (
	"fmt"
	"github.com/abhinandpn/ImageCompress/internal/domain"
)

// ImageUsecase handles the business logic for image compression
type ImageUsecase struct {
	service domain.ImageService
}

// NewImageUsecase creates a new instance of ImageUsecase
func NewImageUsecase(service domain.ImageService) *ImageUsecase {
	return &ImageUsecase{service: service}
}

// CompressImage processes the uploaded image and returns compressed paths
func (u *ImageUsecase) CompressImage(imageData []byte, originalSize int64) (*domain.Image, error) {
	if originalSize > 10*1024*1024 { // Max 10MB
		return nil, fmt.Errorf("image size exceeds 10MB limit")
	}

	paths, err := u.service.CompressImage(imageData, originalSize)
	if err != nil {
		return nil, fmt.Errorf("failed to compress image: %v", err)
	}

	image := domain.NewImage(originalSize)
	image.CompressedPaths = paths
	return image, nil
}