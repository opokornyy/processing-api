# Makefile

.PHONY: build-container run

# Variables for the image name and tag.
IMAGE_NAME := processing-app
IMAGE_TAG  := latest

build-container:
	@echo "Building container image $(IMAGE_NAME):$(IMAGE_TAG)..."
	@podman build -t $(IMAGE_NAME):$(IMAGE_TAG) .

# 'build-container' is now a prerequisite for 'run'
run: build-container
	@echo "Running container $(IMAGE_NAME)..."
	@podman run -d --name $(IMAGE_NAME) --env-file ../.env -p 8080:8080 $(IMAGE_NAME):$(IMAGE_TAG)