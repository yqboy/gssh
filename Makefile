BIN="bin/gssh"
VERSION=V0.0.1

build:
	@env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w -X main.Version=${VERSION}' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ${BIN}_linux main.go
	@env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w -X main.Version=${VERSION}' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ${BIN}_mac main.go
	@env CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags '-s -w -X main.Version=${VERSION}' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ${BIN}_mac_arm main.go
