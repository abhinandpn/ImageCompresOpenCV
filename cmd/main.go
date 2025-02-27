package main

import (
	"log"
	"github.com/labstack/echo/v4"
	"github.com/abhinandpn/ImageCompress/internal/delivery"
	"github.com/abhinandpn/ImageCompress/internal/domain"
	"github.com/abhinandpn/ImageCompress/internal/infrastructure"
	"github.com/abhinandpn/ImageCompress/internal/usecase"
)

// main initializes and starts the HTTP server
func main() {
	e := echo.New()

	// Serve static files from the uploads directory (one level up from cmd)
	e.Static("/uploads", "../uploads")

	// Initialize dependencies
	imaginaryRepo := infrastructure.NewImaginaryRepository()
	imageService := domain.NewImageService(imaginaryRepo)
	imageUsecase := usecase.NewImageUsecase(imageService)
	handler := delivery.NewHTTPHandler(imageUsecase)

	// Register routes
	e.POST("/upload", handler.UploadImage)

	// Start server
	log.Fatal(e.Start(":8080"))
}