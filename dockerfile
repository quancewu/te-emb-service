# ---dev---
FROM golang:1.23 AS dev

WORKDIR /opt/app/te-api

ARG TARGETOS
ARG TARGETARCH
ARG TARGETPLATFORM

ENV MODE=docker-mode

RUN apt update

RUN apt install make

# build te emb service
ADD te-emb-api /opt/app/te-api/te-emb-api

WORKDIR /opt/app/te-api/te-emb-api

RUN go mod download

RUN make build

RUN if [ "$TARGETPLATFORM" = "linux/arm64" ]; then \
make buildarm ;\
elif [ "$TARGETPLATFORM" = "linux/amd64" ]; then \
make build ;\
fi

# build te redis servcie
ADD te-redis-service /opt/app/te-api/te-redis-service

WORKDIR /opt/app/te-api/te-redis-service

RUN go mod download

RUN if [ "$TARGETPLATFORM" = "linux/arm64" ]; then \
make buildarm ;\
elif [ "$TARGETPLATFORM" = "linux/amd64" ]; then \
make build ;\
fi

WORKDIR /opt/app/te-api

# ---release---
FROM --platform=$BUILDPLATFORM golang:1.23-alpine AS release

ARG TARGETOS
ARG TARGETARCH

WORKDIR /opt/app/te-api

ENV MODE=docker-mode

COPY --from=dev /opt/app/te-api/te-redis-service/bin/te-redis-service /opt/app/te-api/

# COPY ./te-emb-api/bin/te-emb-api /opt/app/te-api/

COPY --from=dev /opt/app/te-api/te-emb-api/bin/te-emb-api /opt/app/te-api/

COPY ./script.sh /opt/app/te-api/

RUN chmod +x te-redis-service

RUN chmod +x te-emb-api

RUN chmod +x script.sh

RUN pwd

CMD ["./script.sh"]

