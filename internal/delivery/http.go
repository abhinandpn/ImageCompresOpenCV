package delivery

import (
	"net/http"
	"path/filepath"

	"github.com/abhinandpn/ImageCompress/internal/usecase"
	"github.com/labstack/echo/v4"
)

// HTTPHandler handles HTTP requests for the image compressor
type HTTPHandler struct {
	usecase *usecase.ImageUsecase
}

// NewHTTPHandler creates a new HTTPHandler instance
func NewHTTPHandler(usecase *usecase.ImageUsecase) *HTTPHandler {
	return &HTTPHandler{usecase: usecase}
}

// UploadImage handles image upload and compression
func (h *HTTPHandler) UploadImage(c echo.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "image file is required"})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to open image"})
	}
	defer src.Close()

	// Read image data
	data := make([]byte, file.Size)
	_, err = src.Read(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to read image"})
	}

	// Compress image
	image, err := h.usecase.CompressImage(data, file.Size)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Convert paths to URLs
	baseURL := "http://localhost:8080"
	urls := make(map[string]string)
	for key, path := range image.CompressedPaths {
		// Use only the filename in the URL since /uploads is served statically
		urls[key] = baseURL + "/uploads/" + filepath.Base(path)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"paths": urls,
	})
}
