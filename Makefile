# Путь к директории с .proto файлами
PROTO_DIR := ./proto

# Путь к директории, куда будет генерироваться Go-код
GO_OUT_DIR := ./grpc

# Список .proto файлов
PROTO_FILES := $(PROTO_DIR)/settings.proto $(PROTO_DIR)/automode.proto  $(PROTO_DIR)/log.proto  $(PROTO_DIR)/roles.proto  $(PROTO_DIR)/welcome.proto

# Команда генерации gRPC и Go-кода
PROTOC_COMMAND = protoc -I=$(PROTO_DIR) \
	--go_out=$(GO_OUT_DIR) \
	--go-grpc_out=$(GO_OUT_DIR)

.PHONY: all generate docker-up lint

all: generate

generate-grpc:
	@echo "Generating Go code from .proto files..."
	@mkdir -p $(GO_OUT_DIR)
	@$(PROTOC_COMMAND) $(PROTO_FILES)
	@echo "Go code generated in $(GO_OUT_DIR)"


grpc-clean:
	@echo "Cleaning up generated Go code..."
	@rm -rf $(GO_OUT_DIR)/*
	@echo "Cleaned up generated Go code in $(GO_OUT_DIR)"

docker-run:
	docker compose up --build


grpc-init:
	@echo "Initializing gRPC server..."
	@$(MAKE) generate-grpc

	@echo "Copying all generated files to api-gateway/proto..."
	@mkdir -p api-gateway/proto
	@cp -r grpc/* api-gateway/

	@echo "Moving settings proto files to settings-service..."
	@mkdir -p settings-service/proto
	@cp -r grpc/* settings-service/
	@echo "gRPC server initialized. Files distributed to services."

	@$(MAKE) grpc-clean

lint: 
	gofmt -w bot/ settings-service/ api-gateway/
	@echo "Code formatted with gofmt."

docs: 
	swag init -g api-gateway/main.go --output api-gateway/docs
	@echo "Swagger docs generated."