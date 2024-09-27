# Makefile

# Define the Go command
GO := go
TARGETS := file.go main.go peer.go
OUTPUT := myapp.exe # Change this to your desired executable name

# Default target to run the application
.PHONY: run
run:
	$(GO) run $(TARGETS)

# Target to build the executable
.PHONY: build
build:
	$(GO) build -o $(OUTPUT) $(TARGETS)

# Clean target to remove compiled files
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f $(OUTPUT)
