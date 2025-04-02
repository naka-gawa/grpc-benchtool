# Version
PROTOC_VERSION := 23.4
GO_PROTOC_GEN_VERSION := v1.24
GO_PROTOC_GEN_GRPC_VERSION := v1.1.0
COBRA_VERSION := v1.3.0

# Variables
PROTOC ?= protoc
PROTOC_BIN_DIR := bin
PROTO_DEF_DIR := proto-def
PROTO_DIR := proto
PROTO_FILES := $(wildcard $(PROTO_DEF_DIR)/*/*.proto)

GO_OUT_FLAGS := \
  --go_out=$(PROTO_DIR) \
  --go_opt=paths=source_relative \
  --go_opt=Mgrpcbench/grpcbench.proto=github.com/naka-gawa/grpc-benchtool/proto/grpcbench \
  --go-grpc_out=$(PROTO_DIR) \
  --go-grpc_opt=paths=source_relative \
  --go-grpc_opt=Mgrpcbench/grpcbench.proto=github.com/naka-gawa/grpc-benchtool/proto/grpcbench \
  --proto_path=$(PROTO_DEF_DIR)

.PHONY: all proto clean setup setup.protoc setup.protoc-gen-go setup.protoc-gen-go-grpc setup.cobra

all: proto

proto:
	@echo "Generating Go code from proto files..."
	PATH="$(PWD)/$(PROTOC_BIN_DIR):$$PATH" $(PROTOC) $(GO_OUT_FLAGS) $(PROTO_FILES)

clean:
	@echo "Cleaning generated files..."
	@rm -rf $(PROTO_DIR)/grpcbench/*.pb.go

cobra: add
	@echo "Generating cobra commands..."
	@GOBIN=$(PWD)/$(PROTOC_BIN_DIR) go install github.com/spf13/cobra-cli@$(COBRA_VERSION)
	@cobra-cli init --pkg-name grpcbench --output-dir cmd
	@cobra-cli add grpcbench --output-dir cmd

setup: setup.protoc setup.protoc-gen-go setup.protoc-gen-go-grpc

setup.protoc:
	@echo "Checking for protoc..."
	@command -v $(PROTOC_BIN_DIR)/protoc >/dev/null 2>&1 && { \
		echo "protoc is already installed in $(PROTOC_BIN_DIR)."; \
	} || { \
		echo "protoc not found. Installing..."; \
		if ! command -v curl >/dev/null 2>&1; then \
			echo "Error: curl is required but not installed."; exit 1; \
		fi; \
		if ! command -v unzip >/dev/null 2>&1; then \
			echo "Error: unzip is required but not installed."; exit 1; \
		fi; \
		OS=$$(uname -s | tr A-Z a-z); \
		if [ "$$OS" = "darwin" ]; then OS="osx"; fi; \
		ARCH=$$(uname -m); \
		if [ "$$ARCH" = "x86_64" ]; then ARCH="x86_64"; \
		elif [ "$$ARCH" = "arm64" ]; then ARCH="aarch_64"; \
		else echo "Unsupported architecture: $$ARCH"; exit 1; fi; \
		PROTOC_ZIP=protoc-$(PROTOC_VERSION)-$$OS-$$ARCH.zip; \
		curl -Lo /tmp/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/$$PROTOC_ZIP; \
		unzip -o /tmp/protoc.zip bin/protoc -d ./$(PROTOC_BIN_DIR); \
		rm -f /tmp/protoc.zip; \
		echo "protoc installed in $(PROTOC_BIN_DIR)"; \
		echo "You may want to add this to your PATH:"; \
		echo "  export PATH=\"$$(pwd)/$(PROTOC_BIN_DIR):\$$PATH\""; \
	}

setup.protoc-gen-go:
	@echo "Installing protoc-gen-go..."
	@GOBIN=$(PWD)/$(PROTOC_BIN_DIR) go install google.golang.org/protobuf/cmd/protoc-gen-go@$(GO_PROTOC_GEN_VERSION)

setup.protoc-gen-go-grpc:
	@echo "Installing protoc-gen-go-grpc..."
	@GOBIN=$(PWD)/$(PROTOC_BIN_DIR) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(GO_PROTOC_GEN_GRPC_VERSION)

# To generate a new cobra command:
#   cobra-cli add server
#   cobra-cli add client
setup.cobra:
	@echo "Installing cobra..."
	@GOBIN=$(PWD)/$(PROTOC_BIN_DIR) go install github.com/spf13/cobra-cli@$(COBRA_VERSION)
