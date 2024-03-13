#!/bin/bash

SHELL        := /bin/bash
SERVICE      := rent.car
PROTOS       := car

# Proto generate proto files
.PHONY: proto
proto:
	@go mod vendor
	@protoc  -I ./vendor --proto_path=${GOPATH}/src:. --nvo_out=paths=source_relative:. --nvo_opt=name=$(SERVICE) --go_out=paths=source_relative:. protos/$(PROTOS).proto
	@rm -rf vendor
