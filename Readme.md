ImageCompress - Project Documentation

Overview
--------
ImageCompress is a Go-based API that allows users to upload an image and receive four compressed versions of it in different size ranges. The application uses the bimg library (backed by libvips) for efficient image processing and follows a Clean Architecture design for maintainability and scalability.

Features
--------
- Upload an image (max 10MB).
- Compress the image into four sizes:
  - Original: Adjusted based on size (50% quality for 5-10MB, 70% for 2-5MB, 100% for <2MB).
  - 300KB: 200-300KB.
  - 20KB: 10-20KB.
  - 5KB: 1-5KB.
- Save compressed images in the uploads directory with readable filenames (e.g., 20250227_195029_original.jpg).
- Serve images via HTTP for easy access.
- Return URLs to the compressed images in the API response.

Project Structure
-----------------
ImageCompress/
├── cmd/
│   └── main.go              Entry point of the application
├── internal/
│   ├── domain/              Business logic (entities and services)
│   │   ├── image.go
│   │   └── service.go
│   ├── usecase/             Application logic (use cases)
│   │   └── image_usecase.go
│   ├── infrastructure/      External systems (image processing)
│   │   └── imaginary.go
│   └── delivery/            API handlers (HTTP)
│       └── http.go
├── uploads/                 Directory for compressed images
├── go.mod                   Go module dependencies
└── README.md                Project documentation

Prerequisites
-------------
- Go: Version 1.21 or higher.
- libvips: Version 8.8+ (required by bimg).
- pkg-config: To locate libvips during compilation.

Installation
------------
1. Clone the Repository
   Command: git clone https://github.com/abhinandpn/ImageCompress.git
   Command: cd ImageCompress

2. Install Dependencies
   Initialize the Go module and fetch dependencies:
   Command: go mod init github.com/abhinandpn/ImageCompress
   Command: go get github.com/h2non/bimg
   Command: go get github.com/labstack/echo/v4

3. Install libvips
   On Ubuntu/Debian:
   Command: sudo apt update
   Command: sudo apt install -y libvips-dev pkg-config
   Verify installation:
   Command: vips --version  (Should output 8.8+ e.g., 8.13.3)
   Command: pkg-config --modversion vips
   If pkg-config can't find vips, set the path:
   Command: export PKG_CONFIG_PATH=/usr/lib/x86_64-linux-gnu/pkgconfig:$PKG_CONFIG_PATH

   Alternative (if apt fails):
   Command: curl -s https://raw.githubusercontent.com/h2non/bimg/master/preinstall.sh | sudo bash -

4. Create Uploads Directory
   Command: mkdir -p uploads
   Command: chmod 755 uploads

Running the Application
-----------------------
Run the server from the cmd directory:
Command: cd cmd
Command: go run main.go
- The server starts on http://localhost:8080.
- Images are saved in ../uploads/ (relative to cmd).

Usage
-----
API Endpoint
- Method: POST
- URL: http://localhost:8080/upload
- Request: multipart/form-data
- Body:
  - Key: image
  - Value: [Your image file, max 10MB]
- Response: JSON with URLs to compressed images

Example Request (cURL)
Command: curl -X POST -F "image=@path/to/your/image.jpg" http://localhost:8080/upload

Example Response
{
    "paths": {
        "original": "http://localhost:8080/uploads/20250227_195029_original.jpg",
        "300kb": "http://localhost:8080/uploads/20250227_195029_300kb.jpg",
        "20kb": "http://localhost:8080/uploads/20250227_195029_20kb.jpg",
        "5kb": "http://localhost:8080/uploads/20250227_195029_5kb.jpg"
    }
}

Testing with Postman
1. Open Postman and create a new request.
2. Set Method to POST and URL to http://localhost:8080/upload.
3. Go to the Body tab, select form-data.
4. Add a key image, change type to File, and upload an image.
5. Click Send.
6. Copy a URL from the response (e.g., http://localhost:8080/uploads/20250227_195029_original.jpg), create a new GET request, and view the image.

Viewing Images
- File System: Images are stored in ~/Desktop/ImageCompress/uploads/.
- Browser: Use the URLs from the response (e.g., http://localhost:8080/uploads/20250227_195029_original.jpg).

Compression Logic
-----------------
- Original:
  - >5MB: 50% quality.
  - 2-5MB: 70% quality.
  - <2MB: 100% quality (original).
- 300KB: Targets 200-300KB, adjusts resolution if needed.
- 20KB: Targets 10-20KB.
- 5KB: Targets 1-5KB.
- Fallback: If target size isn't met, uses quality 50 with resolution scaling.

Troubleshooting
---------------
- Permission Denied: Ensure uploads has write permissions:
  Command: chmod 755 ~/Desktop/ImageCompress/uploads
- libvips Not Found: Reinstall libvips and verify pkg-config.
- Size Errors: Check logs in the terminal for compression attempts and fallbacks.

Dependencies
------------
- github.com/h2non/bimg: Image processing library.
- github.com/labstack/echo/v4: HTTP framework.

Contributing
------------
1. Fork the repository.
2. Create a feature branch (git checkout -b feature/your-feature).
3. Commit changes (git commit -m "Add your feature").
4. Push to the branch (git push origin feature/your-feature).
5. Open a pull request.

License
-------
MIT License - Abhinand P N