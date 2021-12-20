APP_NAME = bender-apiserver
BUILD_DIR = $(PWD)/build

clean:
	rm -rf ./build

security:
	gosec -quiet ./...

build: clean security swag
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) .

run: build
	$(BUILD_DIR)/$(APP_NAME)

swag:
	swag init --dir .,controllers/.,controllers/.,models/.
	