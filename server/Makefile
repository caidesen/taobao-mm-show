APP_NAME = tbmm
BUILD_DIR = $(PWD)/build

build: clean
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME)_drawin_arm_64 main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME)_linux_amd_64 main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME)_windows_amd_64 main.go

clean:
	rm -rf ./build

