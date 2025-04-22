# Путь к директории с .proto файлами
PROTO_DIR := ./proto

# Путь к директории, куда будет генерироваться Go-код
GO_OUT_DIR := ./grpc

# Список .proto файлов
PROTO_FILES := $(PROTO_DIR)/gaming.proto $(PROTO_DIR)/settings.proto

# Команда генерации gRPC и Go-кода
PROTOC_COMMAND = protoc -I=$(PROTO_DIR) \
	--go_out=$(GO_OUT_DIR) \
	--go-grpc_out=$(GO_OUT_DIR)

.PHONY: all generate docker-up lint

all: generate

generate:
	@echo "Generating Go code from .proto files..."
	@mkdir -p $(GO_OUT_DIR)
	@$(PROTOC_COMMAND) $(PROTO_FILES)
	@echo "Go code generated in $(GO_OUT_DIR)"


grpc-clean:
	@echo "Cleaning up generated Go code..."
	@rm -rf $(GO_OUT_DIR)/*
	@echo "Cleaned up generated Go code in $(GO_OUT_DIR)"

docker:
	docker compose up --build


grpc-init:
	@echo "Initializing gRPC server..."
	@$(MAKE) generate

	@echo "Copying all generated files to api-gateway/proto..."
	@mkdir -p api-gateway/proto
	@cp -r grpc/* api-gateway/proto/

	@echo "Moving settings proto files to settings-service..."
	@mkdir -p settings-service/proto
	@find grpc -type f -name "settings*.pb.go" -exec mv {} settings-service/proto/ \;

	@echo "Moving gaming proto files to gachas-service..."
	@mkdir -p gaming-service/proto
	@find grpc -type f -name "gaming*.pb.go" -exec mv {} gachas-service/proto/ \;

	@rm -rf grpc

	@echo "gRPC server initialized. Files distributed to services."

lint: 
	gofmt -w bot/ settings-service/ api-gateway/ gachas-service/