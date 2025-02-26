BINARY_NAME=social-api
# Set the default configuration file if not provided
CONFIG_FILE=./resource/config/config.yml

# Default target to build and run the Go program
run:
	go run main.go --config=${CONFIG_FILE}

# Target to build the Go program
build:
	go build -o ${BINARY_NAME} main.go

# Clean up the built binary
clean:
	rm -f  ${BINARY_NAME}
