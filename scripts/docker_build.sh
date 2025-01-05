#!/bin/bash
docker run --privileged --rm tonistiigi/binfmt --install all

docker buildx build -t harbor.synclytic.app/te-emb-api/te-emb-service:1.0-alpine --platform linux/arm64/v8,linux/amd64 .

docker push harbor.synclytic.app/te-emb-api/te-emb-service:1.0-alpine
