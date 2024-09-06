# Makefile for spinning up the development environment

# Define variables
UMBRACO_APP_DIR := ./umbraco-app

# Default target
.PHONY: dev
dev:
	@echo "Starting development environment in $(UMBRACO_APP_DIR)..."
	cd $(UMBRACO_APP_DIR) && npm install && npm run dev

# Optional cleanup target
.PHONY: clean
clean:
	@echo "Cleaning node_modules in $(UMBRACO_APP_DIR)..."
	cd $(UMBRACO_APP_DIR) && rm -rf node_modules

# Target to install dependencies
.PHONY: install
install:
	@echo "Installing dependencies in $(UMBRACO_APP_DIR)..."
	cd $(UMBRACO_APP_DIR) && npm install

# Run build
.PHONY: build
build:
	@echo "Building project in $(UMBRACO_APP_DIR)..."
	cd $(UMBRACO_APP_DIR) && npm run build
