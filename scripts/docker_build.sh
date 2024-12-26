#!/bin/bash
docker run --privileged --rm tonistiigi/binfmt --install all

docker buildx build -t harbor.syncltic.app/te-emb-api/te-emb-service:alpha-0.9 --platform linux/arm64/v8,linux/amd64 .

docker push harbor.syncltic.app/te-emb-api/te-emb-service:alpha-0.9
