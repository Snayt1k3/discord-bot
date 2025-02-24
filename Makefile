PROTO_DIR=proto
OUT_DIR=settings-service/proto
GATEWAY_OUT_DIR=api-gateway/proto
PROTO_FILES=$(wildcard $(PROTO_DIR)/*.proto)

.PHONY: all generate copy clean

all: generate copy

generate:
	@echo "Generating gRPC code..."
	mkdir -p $(OUT_DIR)
	protoc --proto_path=$(PROTO_DIR) \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)

copy:
	@echo "Copying generated files to api-gateway..."
	mkdir -p $(GATEWAY_OUT_DIR)
	cp -r $(OUT_DIR)/* $(GATEWAY_OUT_DIR)/

	@echo "Files copied successfully."

clean:
	@echo "Cleaning generated files..."
	rm -rf $(OUT_DIR) $(GATEWAY_OUT_DIR)
	@echo "Cleanup complete."

docker-up:
	docker compose up --build