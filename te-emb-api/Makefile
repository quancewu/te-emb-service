buildarm:
	@env GOOS=linux GOARCH=arm CC=arm-linux-gnueabihf-gcc CXX=arm-linux-gnueabihf-g++ CGO_ENABLED=1 go build -a -ldflags '-extldflags "-static"'  -o "bin/te-emb-api-arm"

build:
	@go build -a -ldflags '-extldflags "-static"'  -o bin/te-emb-api

run: build
	@./bin/te-emb-api -online

test:
	@go test -v ./...