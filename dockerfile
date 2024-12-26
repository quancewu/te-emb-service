# ---dev---
FROM --platform=$BUILDPLATFORM golang:1.23 AS build

WORKDIR /opt/app/te-api

ARG TARGETOS TARGETARCH TARGETPLATFORM

ENV MODE=docker-mode

RUN apt update

RUN apt install gcc-arm-linux-gnueabihf -y

# build te emb service
ADD te-emb-api /opt/app/te-api/te-emb-api

WORKDIR /opt/app/te-api/te-emb-api

RUN go mod download

RUN if [ "$TARGETPLATFORM" = "linux/arm64" ]; then \
GOOS=$TARGETOS GOARCH=arm CC=arm-linux-gnueabihf-gcc CXX=arm-linux-gnueabihf-g++ CGO_ENABLED=1 go build -a -ldflags '-extldflags "-static"'  -o bin/te-emb-api ;\
elif [ "$TARGETPLATFORM" = "linux/amd64" ]; then \
GOOS=$TARGETOS GOARCH=$TARGETARCH go build -a -ldflags '-extldflags "-static"'  -o bin/te-emb-api ;\
fi

# build te redis servcie
ADD te-redis-service /opt/app/te-api/te-redis-service

WORKDIR /opt/app/te-api/te-redis-service

RUN go mod download

RUN if [ "$TARGETPLATFORM" = "linux/arm64" ]; then \
GOOS=$TARGETOS GOARCH=arm CC=arm-linux-gnueabihf-gcc CXX=arm-linux-gnueabihf-g++ CGO_ENABLED=1 go build -a -ldflags '-extldflags "-static"'  -o bin/te-redis-service ;\
elif [ "$TARGETPLATFORM" = "linux/amd64" ]; then \
GOOS=$TARGETOS GOARCH=$TARGETARCH go build -a -ldflags '-extldflags "-static"'  -o bin/te-redis-service ;\
fi

WORKDIR /opt/app/te-api

# ---release---
FROM alpine:3.21 AS release

ARG TARGETOS
ARG TARGETARCH

WORKDIR /opt/app/te-api

ENV MODE=docker-mode

COPY --from=build /opt/app/te-api/te-redis-service/bin/te-redis-service /opt/app/te-api/

COPY --from=build /opt/app/te-api/te-emb-api/bin/te-emb-api /opt/app/te-api/

COPY ./script.sh /opt/app/te-api/

RUN chmod +x te-redis-service

RUN chmod +x te-emb-api

RUN chmod +x script.sh

RUN pwd

CMD ["./script.sh"]

