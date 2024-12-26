<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://github.com/quancewu/te-emb-service/blob/main/picture/favicon.svg" alt="Project logo"></a>
</p>

<h3 align="center">te-emb-service</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/quancewu/te-emb-service.svg)](https://github.com/quancewu/te-emb-service/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/quancewu/te-emb-service.svg)](https://github.com/quancewu/te-emb-service/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

<p align="center"> TornadoEdge Embedded System provide redis pub/sub sqlite mgmt
    <br>
</p>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Deployment](#deployment)
- [Usage](#usage)
- [Built Using](#built_using)
- [TODO](../TODO.md)
- [Contributing](../CONTRIBUTING.md)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)

## 🧐 About <a name = "about"></a>

te embedded service handling redis sub to SQLite DB provide Gin REST API

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) for notes on how to deploy the project on a live system.

### Prerequisites

Get Linux golang devlop envirments

```
example@server$ go version
go version go1.23.2 linux/amd64
```

### Installing

Install docker on linux system

Build Docker

```
docker compose up -build -d
```

```
docker buildx build -t harbor.syncltic.app/te-emb-api/te-emb-service:alpha-0.9 --platform linux/arm64/v8,linux/amd64 .
```

## 🔧 Running the tests <a name = "tests"></a>

check docker run

```
executing task: docker compose -f "docker-compose.yaml" up -d --build 
root@example-server:~/project/te-emb-service$ docker buildx build -t harbor.syncltic.app/te-emb-api/te-emb-service:alpha-0.9 --platform linux/arm64/v8,linux/amd64 . --no-cache
[+] Building 153.1s (49/49) FINISHED                                                                                               docker:default
 => [internal] load build definition from dockerfile                                                                                         0.0s
 => => transferring dockerfile: 1.73kB                                                                                                       0.0s
 => [linux/amd64 internal] load metadata for docker.io/library/alpine:3.21                                                                   0.8s
 => [linux/amd64 internal] load metadata for docker.io/library/golang:1.23                                                                   0.9s
 => [linux/arm64 internal] load metadata for docker.io/library/alpine:3.21                                                                   0.9s
 => [internal] load .dockerignore                                                                                                            0.0s
 => => transferring context: 2B                                                                                                              0.0s
 => [linux/arm64 release 1/9] FROM docker.io/library/alpine:3.21@sha256:21dc6063fd678b478f57c0e13f47560d0ea4eeba26dfc947b2a4f81f686b9f45     0.1s
 => => resolve docker.io/library/alpine:3.21@sha256:21dc6063fd678b478f57c0e13f47560d0ea4eeba26dfc947b2a4f81f686b9f45                         0.1s
 => [internal] load build context                                                                                                            0.0s
 => => transferring context: 2.17kB                                                                                                          0.0s
 => [linux/amd64 build  1/13] FROM docker.io/library/golang:1.23@sha256:7ea4c9dcb2b97ff8ee80a67db3d44f98c8ffa0d191399197007d8459c1453041     0.1s
 => => resolve docker.io/library/golang:1.23@sha256:7ea4c9dcb2b97ff8ee80a67db3d44f98c8ffa0d191399197007d8459c1453041                         0.1s
 => [linux/amd64 release 1/9] FROM docker.io/library/alpine:3.21@sha256:21dc6063fd678b478f57c0e13f47560d0ea4eeba26dfc947b2a4f81f686b9f45     0.1s
 => => resolve docker.io/library/alpine:3.21@sha256:21dc6063fd678b478f57c0e13f47560d0ea4eeba26dfc947b2a4f81f686b9f45                         0.1s
 => CACHED [linux/arm64 release 2/9] WORKDIR /opt/app/te-api                                                                                 0.0s
 => CACHED [linux/amd64 build  2/13] WORKDIR /opt/app/te-api                                                                                 0.0s
 => [linux/amd64 build  3/13] RUN apt update                                                                                                 4.4s
 => [linux/amd64->arm64 build  3/13] RUN apt update                                                                                          3.6s
 => CACHED [linux/amd64 release 2/9] WORKDIR /opt/app/te-api                                                                                 0.0s
 => [linux/amd64->arm64 build  4/13] RUN apt install gcc-arm-linux-gnueabihf -y                                                              8.3s
 => [linux/amd64 build  4/13] RUN apt install gcc-arm-linux-gnueabihf -y                                                                     8.7s
 => [linux/amd64->arm64 build  5/13] ADD te-emb-api /opt/app/te-api/te-emb-api                                                               0.3s
 => [linux/amd64->arm64 build  6/13] WORKDIR /opt/app/te-api/te-emb-api                                                                      0.1s
 => [linux/amd64->arm64 build  7/13] RUN go mod download                                                                                     5.4s
 => [linux/amd64 build  5/13] ADD te-emb-api /opt/app/te-api/te-emb-api                                                                      0.3s
 => [linux/amd64 build  6/13] WORKDIR /opt/app/te-api/te-emb-api                                                                             0.1s
 => [linux/amd64 build  7/13] RUN go mod download                                                                                            4.2s
 => [linux/amd64 build  8/13] RUN if [ "linux/amd64" = "linux/arm64" ]; then GOOS=linux GOARCH=arm CC=arm-linux-gnueabihf-gcc CXX=arm-linu  69.7s
 => [linux/amd64->arm64 build  8/13] RUN if [ "linux/arm64" = "linux/arm64" ]; then GOOS=linux GOARCH=arm CC=arm-linux-gnueabihf-gcc CXX=a  70.7s
 => [linux/amd64 build  9/13] ADD te-redis-service /opt/app/te-api/te-redis-service                                                          0.3s
 => [linux/amd64 build 10/13] WORKDIR /opt/app/te-api/te-redis-service                                                                       0.1s
 => [linux/amd64 build 11/13] RUN go mod download                                                                                            2.3s
 => [linux/amd64->arm64 build  9/13] ADD te-redis-service /opt/app/te-api/te-redis-service                                                   0.2s
 => [linux/amd64->arm64 build 10/13] WORKDIR /opt/app/te-api/te-redis-service                                                                0.1s
 => [linux/amd64->arm64 build 11/13] RUN go mod download                                                                                     2.5s
 => [linux/amd64 build 12/13] RUN if [ "linux/amd64" = "linux/arm64" ]; then GOOS=linux GOARCH=arm CC=arm-linux-gnueabihf-gcc CXX=arm-linu  54.5s
 => [linux/amd64->arm64 build 12/13] RUN if [ "linux/arm64" = "linux/arm64" ]; then GOOS=linux GOARCH=arm CC=arm-linux-gnueabihf-gcc CXX=a  55.9s
 => [linux/amd64 build 13/13] WORKDIR /opt/app/te-api                                                                                        0.1s
 => [linux/amd64 release 3/9] COPY --from=build /opt/app/te-api/te-redis-service/bin/te-redis-service /opt/app/te-api/                       0.1s
 => [linux/amd64 release 4/9] COPY --from=build /opt/app/te-api/te-emb-api/bin/te-emb-api /opt/app/te-api/                                   0.2s
 => [linux/amd64 release 5/9] COPY ./script.sh /opt/app/te-api/                                                                              0.1s
 => [linux/amd64 release 6/9] RUN chmod +x te-redis-service                                                                                  0.4s
 => [linux/amd64 release 7/9] RUN chmod +x te-emb-api                                                                                        0.4s
 => [linux/amd64 release 8/9] RUN chmod +x script.sh                                                                                         0.4s
 => [linux/amd64 release 9/9] RUN pwd                                                                                                        0.3s
 => [linux/amd64->arm64 build 13/13] WORKDIR /opt/app/te-api                                                                                 0.1s
 => [linux/arm64 release 3/9] COPY --from=build /opt/app/te-api/te-redis-service/bin/te-redis-service /opt/app/te-api/                       0.1s
 => [linux/arm64 release 4/9] COPY --from=build /opt/app/te-api/te-emb-api/bin/te-emb-api /opt/app/te-api/                                   0.1s
 => [linux/arm64 release 5/9] COPY ./script.sh /opt/app/te-api/                                                                              0.1s
 => [linux/arm64 release 6/9] RUN chmod +x te-redis-service                                                                                  0.3s
 => [linux/arm64 release 7/9] RUN chmod +x te-emb-api                                                                                        0.4s
 => [linux/arm64 release 8/9] RUN chmod +x script.sh                                                                                         0.3s
 => [linux/arm64 release 9/9] RUN pwd                                                                                                        0.3s
 => exporting to image                                                                                                                       2.5s
 => => exporting layers                                                                                                                      1.8s
 => => exporting manifest sha256:6683432cb3cc9a091ac4858858d356ac095a60b271ab68dfaa1edc862e158258                                            0.0s
 => => exporting config sha256:7beebd5d5a06e0f94e8cf4c966dc32db51629a1540846c7543f79a1555998da5                                              0.0s
 => => exporting attestation manifest sha256:19afb2b433385a9b66a1caca817473f85ffeee1862a8081045cb661407aa4f97                                0.0s
 => => exporting manifest sha256:c8a69d2b716221d8d320015e9bfbfa7a1c54b2d30a7b39e0a82bf95d04ac6935                                            0.0s
 => => exporting config sha256:e63bd5b3ab8bb59fc654cf2850f80b0b6c71c0bd3d1dd718ff7e27e92eb3ab6f                                              0.0s
 => => exporting attestation manifest sha256:6908ae5d2500c7fe5c48a90d1c8dadc5a6f218f91c7711b78c62530d8b1c49b8                                0.0s
 => => exporting manifest list sha256:7db7a617adc2b62004208aac2034369eea610c2ca50c4181b97c5ab9ca05cc06                                       0.0s
 => => naming to harbor.syncltic.app/te-emb-api/te-emb-service:alpha-0.9                                                                     0.0s
 => => unpacking to harbor.syncltic.app/te-emb-api/te-emb-service:alpha-0.9                                                                  0.4s
```

### Break down into end to end tests


### And coding style tests


## 🎈 Usage <a name="usage"></a>

Add notes about how to use the system.

## 🚀 Deployment <a name = "deployment"></a>

Add additional notes about how to deploy this on a live system.

## ⛏️ Built Using <a name = "built_using"></a>

## ✍️ Authors <a name = "authors"></a>

- [@quancewu](https://github.com/quancewu) - Idea & Initial work

## 🎉 Acknowledgements <a name = "acknowledgement"></a>
