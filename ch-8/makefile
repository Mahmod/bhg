# Main Makefile for handling multiple subprojects

# List of all subdirectories
SUBDIRS := $(wildcard */.)

# Phony targets for build, clean, and help
.PHONY: all build clean help

# Default target
all: build

# Build all subprojects
build:
	@for dir in $(SUBDIRS); do \
		echo "Building $$dir"; \
		$(MAKE) -C $$dir build || exit 1; \
	done

# Clean all subprojects
clean:
	@for dir in $(SUBDIRS); do \
		echo "Cleaning $$dir"; \
		$(MAKE) -C $$dir clean || exit 1; \
	done

# Help command to display makefile usage
help:
	@echo "Usage: make [target]"
	@echo "Available targets:"
	@echo "  all    : Build all subprojects (default target)"
	@echo "  build  : Build all subprojects"
	@echo "  clean  : Clean all subprojects"
	@echo "  help   : Display this help"
