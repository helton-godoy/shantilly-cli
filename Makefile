.PHONY: build test clean install build-all

# Build for current platform
build:
	@echo "Building Shantilly-CLI..."
	go build -o bin/shantilly cmd/shantilly/main.go
	@echo "✅ Build complete: bin/shantilly"

# Cross-platform builds
build-all:
	@echo "Building for all platforms..."
	@mkdir -p bin
	
	# Linux AMD64
	GOOS=linux GOARCH=amd64 go build -o bin/shantilly-linux-amd64 cmd/shantilly/main.go
	@echo "✅ Linux AMD64: bin/shantilly-linux-amd64"
	
	# Linux ARM64
	GOOS=linux GOARCH=arm64 go build -o bin/shantilly-linux-arm64 cmd/shantilly/main.go
	@echo "✅ Linux ARM64: bin/shantilly-linux-arm64"
	
	# macOS AMD64
	GOOS=darwin GOARCH=amd64 go build -o bin/shantilly-darwin-amd64 cmd/shantilly/main.go
	@echo "✅ macOS AMD64: bin/shantilly-darwin-amd64"
	
	# macOS ARM64 (Apple Silicon)
	GOOS=darwin GOARCH=arm64 go build -o bin/shantilly-darwin-arm64 cmd/shantilly/main.go
	@echo "✅ macOS ARM64: bin/shantilly-darwin-arm64"
	
	# Windows AMD64
	GOOS=windows GOARCH=amd64 go build -o bin/shantilly-windows-amd64.exe cmd/shantilly/main.go
	@echo "✅ Windows AMD64: bin/shantilly-windows-amd64.exe"
	
	@echo "🎉 All builds complete!"

# Run tests
test:
	@echo "Running tests..."
	go test ./...
	@echo "✅ All tests passed!"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	@echo "✅ Clean complete!"

# Install locally
install: build
	@echo "Installing Shantilly-CLI..."
	cp bin/shantilly /usr/local/bin/
	@echo "✅ Installation complete!"

# Install to custom directory
install-custom: build
	@if [ -z "$(DESTDIR)" ]; then \
		echo "Usage: make install-custom DESTDIR=/path/to/install"; \
		exit 1; \
	fi
	@echo "Installing to $(DESTDIR)..."
	cp bin/shantilly $(DESTDIR)/
	@echo "✅ Installation complete!"

# Development build with race detection
dev:
	@echo "Building with race detection..."
	go build -race -o bin/shantilly-dev cmd/shantilly/main.go
	@echo "✅ Dev build complete: bin/shantilly-dev"

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...
	@echo "✅ Code formatted!"

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run
	@echo "✅ Linting complete!"

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy
	@echo "✅ Dependencies updated!"

# Show build info
info:
	@echo "Shantilly-CLI Build Info"
	@echo "======================="
	@echo "Go Version: $(shell go version)"
	@echo "Module: $(shell go list -m)"
	@echo "Git Branch: $(shell git branch --show-current)"
	@echo "Git Commit: $(shell git rev-parse --short HEAD)"

# Quick test build
test-build:
	@echo "Quick test build..."
	go build -o /tmp/shantilly-test cmd/shantilly/main.go
	@echo "✅ Test build successful!"

# Release build (optimized)
release: clean
	@echo "Building release version..."
	@mkdir -p bin
	go build -ldflags="-s -w" -o bin/shantilly cmd/shantilly/main.go
	@echo "✅ Release build complete: bin/shantilly"
	@echo "Binary size: $$(du -h bin/shantilly | cut -f1)"

# Help
help:
	@echo "Shantilly-CLI Makefile"
	@echo "====================="
	@echo ""
	@echo "Available targets:"
	@echo "  build        - Build for current platform"
	@echo "  build-all    - Build for all platforms"
	@echo "  test         - Run tests"
	@echo "  clean        - Clean build artifacts"
	@echo "  install      - Install to /usr/local/bin"
	@echo "  dev          - Development build with race detection"
	@echo "  fmt          - Format code"
	@echo "  lint         - Run linter"
	@echo "  deps         - Download dependencies"
	@echo "  info         - Show build info"
	@echo "  test-build   - Quick test build"
	@echo "  release      - Optimized release build"
	@echo "  help         - Show this help"
