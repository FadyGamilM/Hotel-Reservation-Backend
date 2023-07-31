BINARY_NAME=hotelreservation
build:
	@echo "**** Building the binary of the backend service... ****"
	go build -o bin/${BINARY_NAME} cmd/main.go
	@echo "**** Application building is succeed ! ****"

run:
	./bin/${BINARY_NAME}

test:
	go test -v ./...