VERSION ?= $(shell cat VERSION)

build:
	GOOS=linux GOARCH=amd64 go build -o builds/proftpd_exporter-${VERSION}-linux .
	go build -o builds/proftpd_exporter-${VERSION}-darwin .
