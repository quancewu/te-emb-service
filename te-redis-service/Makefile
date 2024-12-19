buildarm:
	@env GOOS=linux GOARCH=arm CC=arm-linux-gnueabihf-gcc CXX=arm-linux-gnueabihf-g++ CGO_ENABLED=1 go build -a -ldflags '-extldflags "-static"'  -o "bin/te-redis-service-arm"

build:
	@go build -a -ldflags '-extldflags "-static"'  -o bin/te-redis-service

run: build
	@./bin/te-redis-service -online

test:
	@go test -v ./...