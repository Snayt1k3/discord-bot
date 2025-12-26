# Путь к директории с .proto файлами
PROTO_DIR := ./proto

# Список .proto файлов
PROTO_FILES := $(PROTO_DIR)/settings.proto $(PROTO_DIR)/automode.proto  $(PROTO_DIR)/log.proto  $(PROTO_DIR)/roles.proto  $(PROTO_DIR)/welcome.proto  $(PROTO_DIR)/interaction.proto

# Команда генерации gRPC и Go-кода
PROTOC_COMMAND = protoc -I=$(PROTO_DIR) \
	--go_out=$(GO_OUT_DIR) \
	--go-grpc_out=$(GO_OUT_DIR)

.PHONY: all generate docker-up lint

all: generate

generate-grpc:
	@echo "Generating Go code from .proto files..."
	@mkdir -p ./grpc/interaction
	@protoc -I=./proto/interaction \
		--go_out=./grpc/interaction \
		--go-grpc_out=./grpc/interaction \
		./proto/interaction/*.proto
	@mkdir -p "./grpc/preferences"
	@protoc -I=./proto/preferences \
		--go_out=./grpc/preferences \
		--go-grpc_out=./grpc/preferences \
		./proto/preferences/*.proto
	@echo "Go code generated in ./grpc"

interaction-init:
	@mkdir -p "./interaction-service/proto"
	@rm -f ./interaction-service/proto/*.go
	@cp ./grpc/interaction/proto/* ./interaction-service/proto

preferences-init:
	@mkdir -p "./settings-service/proto"
	@rm -f ./settings-service/proto/*.go
	@cp ./grpc/preferences/proto/* ./settings-service/proto

api-gateway-init:
	@mkdir -p "./api-gateway/proto"
	@rm -f ./api-gateway/proto/*.go
	@cp ./grpc/interaction/proto/* ./api-gateway/proto
	@cp ./grpc/preferences/proto/* ./api-gateway/proto

grpc-clean:
	@echo "Cleaning up generated Go code..."
	@rm -rf ./grpc
	@echo "Cleaned up generated Go code in ./grpc"

run:
	docker compose up --build


grpc-init:
	@echo "Initializing gRPC server..."
	@$(MAKE) generate-grpc

	@echo "Copying all generated files to api-gateway/proto..."
	@$(MAKE) api-gateway-init

	@echo "Moving settings proto files to settings-service..."
	@$(MAKE) preferences-init

	@echo "Moving settings proto files to interaction-service..."
	@$(MAKE) interaction-init

	@echo "gRPC server initialized. Files distributed to services."

	@$(MAKE) grpc-clean

lint: 
	gofmt -w bot/ settings-service/ api-gateway/ interaction-service/
	@echo "Code formatted with gofmt."

docs: 
	swag init -g api-gateway/main.go --output api-gateway/docs
	@echo "Swagger docs generated."