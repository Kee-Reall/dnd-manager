BINARY_NAME=dnd-manager
OUT_DIR=./bin
MAIN_PATH=./cmd/main.go
SCPIRT_PATH=scripts


build: clean
	mkdir -p $(OUT_DIR)
	CGO_ENABLED=0 GOOS=linux go build -o $(OUT_DIR)/$(BINARY_NAME) $(MAIN_PATH)

start: build
	$(OUT_DIR)/$(BINARY_NAME)

clean:
	rm -rf $(OUT_DIR)/*

run:
	go run $(MAIN_PATH)

.PHONY: migrategen

migrategen:
	@echo "Usage: make migrategen_<migration_name>"
	@echo "Example: make migrategen_add_index_for_user_table"

migrategen_%:
	python3 $(SCPIRT_PATH)/migrate.py $*