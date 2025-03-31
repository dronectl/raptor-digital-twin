
PROTO_DIR := protocols/protobuf
GEN_DIR := gen

.PHONY: proto
proto:
	@mkdir -p $(GEN_DIR)
	protoc --proto_path=$(PROTO_DIR) \
		--go_out=$(GEN_DIR) --go_opt=paths=source_relative \
		--go_opt=Mraptor/v1/uhci.proto=github.com/dronectl/rdt/$(GEN_DIR)/raptor/v1 \
		--go_opt=Mraptor/v1/commands.proto=github.com/dronectl/rdt/$(GEN_DIR)/raptor/v1 \
		--go_opt=Mraptor/v1/device.proto=github.com/dronectl/rdt/$(GEN_DIR)/raptor/v1 \
		--go_opt=Mraptor/v1/battery.proto=github.com/dronectl/rdt/$(GEN_DIR)/raptor/v1 \
		--go_opt=Mraptor/v1/env.proto=github.com/dronectl/rdt/$(GEN_DIR)/raptor/v1 \
		--go_opt=Mraptor/v1/signals.proto=github.com/dronectl/rdt/$(GEN_DIR)/raptor/v1 \
		--go_opt=Mraptor/v1/profile.proto=github.com/dronectl/rdt/$(GEN_DIR)/raptor/v1 \
		--go_opt=Mraptor/v1/registers.proto=github.com/dronectl/rdt/$(GEN_DIR)/raptor/v1 \
		--go_opt=Mraptor/v1/relay.proto=github.com/dronectl/rdt/$(GEN_DIR)/raptor/v1 \
		--go_opt=Mraptor/v1/stream.proto=github.com/dronectl/rdt/$(GEN_DIR)/raptor/v1 \
		$(PROTO_DIR)/raptor/v1/*.proto


.PHONY: fmt
fmt:
	gofmt -s -w .

