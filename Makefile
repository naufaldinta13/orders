#!/bin/bash

SHELL        := /bin/bash
SERVICE      := testing.car
PROTOS       := car
NEXT_VERSION := $(shell ./version.sh $(dir) $(v))

# Releser function to create new tag version
# based on directory showing the the dependency and push to git
define releaser
	@echo -e "Releasing $(1) version to \033[92m$(2)\033[0m"
	@echo -e "\r\nPlease review the Dependency first:"
	@cd $(1) && go mod graph | grep "env-io" && cd ../
	@echo -e -n "\r\nAre you sure want to releasing it? [y/N] " && read ans && [ $${ans:-N} = y ]
	@git tag -a $(2) -m $(2)
	@echo "Push $(1) version $(2)"
	@git push origin $(2) -q && echo -e "\r\n\033[92m$(SERVICE)@$(2)\033[0m Released!"
endef

# Updater function to update the go mod dependency based on directory
define updater
	@(echo Updating Dependency $(1) && cd $(1) && \
	    go mod tidy && cd ../ );
	@git status | grep $(1)
endef

# Proto generate proto files
.PHONY: proto
proto:
	@go mod vendor
	@protoc  -I ./vendor --proto_path=${GOPATH}/src:. --nvo_out=paths=source_relative:. --nvo_opt=name=$(SERVICE) --go_out=paths=source_relative:. protos/$(PROTOS).proto
	@rm -rf vendor

# Gendoc generate readme rest documentation from go handler
.PHONY: gendoc
gendoc:
	@go get -u go.uber.org/zap@v1.21.0
	@swag init --parseDependency --parseInternal --parseDepth 2 && rm -rf ./docs/docs.go && rm -rf ./docs/swagger.yaml
	@swagger-markdown -i docs/swagger.json -o README.md
	@git checkout go.*

# check-release before releasing the repository should be clean
.PHONY: check-release
check-release:
ifeq ($(v),)
	@(echo -e "Please specify the release type: \n\033[92mv=patch || v=minor || v=major\033[0m\n"; exit 1)
endif
ifeq ($(dir),)
	@(echo -e "Please specify the directory: \n\033[92mdir=. || dir=entity || dir=protos\033[0m\n"; exit 1)
endif
ifneq ($(shell git diff-index --quiet HEAD; echo $$?),0)
	@(echo -e "\033[31mYour code is dirty!\033[0m\nPlease commit the changes first and push it before release!\n"; exit 1)
endif

# Uodate running updater go mod dependency
.PHONY: update
update:
	@$(call updater, $(dir))

# Release running releaser based on directory
.PHONY: release
release: check-release
	@$(call releaser,$(dir),$(NEXT_VERSION))

# Print the current version of repository
.PHONY: version
version:
	@(echo -e "\033[92mService Version :\033[0m " $(shell ./version.sh . current))
	@(echo -e "\033[92mEntity Version  :\033[0m " $(shell ./version.sh entity current))
	@(echo -e "\033[92mProtos Version  :\033[0m " $(shell ./version.sh protos current))
