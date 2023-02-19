API_PATH=api/music
PROTO_OUT_DIR=pkg/music-api
PROTO_API_DIR=$(API_PATH)

.PHONY: gen
gen: gen-proto generate


.PHONY: gen-proto
gen-proto:
	mkdir -p $(PROTO_OUT_DIR)
	protoc \
		-I $(API_PATH)/v1 \
		-I third_party/googleapis \
		--go_out=$(PROTO_OUT_DIR) --go_opt=paths=source_relative \
        ./$(PROTO_API_DIR)/v1/*.proto

generate:
	go generate ./...
run:
	go run cmd/sb-music/main.go