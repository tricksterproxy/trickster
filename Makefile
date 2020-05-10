# Copyright 2018 Comcast Cable Communications Management, LLC
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

DEFAULT: build

PROJECT_DIR    := $(shell pwd)
GO             ?= go
GOFMT          ?= $(GO)fmt
FIRST_GOPATH   := $(firstword $(subst :, ,$(shell $(GO) env GOPATH)))
TRICKSTER_MAIN := cmd/trickster
TRICKSTER      := $(FIRST_GOPATH)/bin/trickster
PROGVER        := $(shell grep 'applicationVersion = ' $(TRICKSTER_MAIN)/main.go | awk '{print $$3}' | sed -e 's/\"//g')
BUILD_TIME     := $(shell date -u +%FT%T%z)
GIT_LATEST_COMMIT_ID     := $(shell git rev-parse HEAD)
GO_VER         := $(shell go version | awk '{print $$3}')
IMAGE_TAG      ?= latest
IMAGE_ARCH     ?= amd64
GOARCH         ?= amd64
TAGVER         ?= unspecified
LDFLAGS         =-ldflags "-extldflags '-static' -w -s -X main.applicationBuildTime=$(BUILD_TIME) -X main.applicationGitCommitID=$(GIT_LATEST_COMMIT_ID) -X main.applicationGoVersion=$(GO_VER) -X main.applicationGoArch=$(GOARCH)"
BUILD_SUBDIR   := OPATH
PACKAGE_DIR    := ./$(BUILD_SUBDIR)/trickster-$(PROGVER)
BIN_DIR        := $(PACKAGE_DIR)/bin
CONF_DIR       := $(PACKAGE_DIR)/conf

.PHONY: validate-app-version
validate-app-version:
	@if [ "$(PROGVER)" != $(TAGVER) ]; then\
		(echo "mismatch between TAGVER '$(TAGVER)' and applicationVersion '$(PROGVER)'"; exit 1);\
	fi

.PHONY: go-mod-vendor
go-mod-vendor:
	$(GO) mod vendor

.PHONY: go-mod-tidy
go-mod-tidy:
	$(GO) mod tidy

.PHONY: test-go-mod
test-go-mod:
	@git diff --quiet --exit-code go.mod go.sum || echo "There are changes to go.mod and go.sum which needs to be committed"

.PHONY: build
build: go-mod-tidy go-mod-vendor
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) $(GO) build $(LDFLAGS) -o ./$(BUILD_SUBDIR)/trickster -a -v $(TRICKSTER_MAIN)/*.go

rpm: build
	mkdir -p ./$(BUILD_SUBDIR)/SOURCES
	cp -p ./$(BUILD_SUBDIR)/trickster ./$(BUILD_SUBDIR)/SOURCES/
	cp $(TRICKSTER_MAIN)/conf/trickster.service ./$(BUILD_SUBDIR)/SOURCES/
	sed -e 's%^# log_file =.*$$%log_file = "/var/log/trickster/trickster.log"%' \
		-e 's%prometheus:9090%localhost:9090%' \
		< $(TRICKSTER_MAIN)/conf/example.conf > ./$(BUILD_SUBDIR)/SOURCES/trickster.conf
	rpmbuild --define "_topdir $(CURDIR)/$(BUILD_SUBDIR)" \
		--define "_version $(PROGVER)" \
		--define "_release 1" \
		-ba deploy/packaging/trickster.spec

.PHONY: install
install:
	$(GO) install -o $(TRICKSTER) $(PROGVER)

.PHONY: release
release: validate-app-version clean go-mod-tidy go-mod-vendor release-artifacts

.PHONY: release-artifacts
release-artifacts: clean

	mkdir -p $(PACKAGE_DIR)
	mkdir -p $(BIN_DIR)
	mkdir -p $(CONF_DIR)

	cp -r ./docs $(PACKAGE_DIR)
	cp -r ./deploy $(PACKAGE_DIR)
	cp ./README.md $(PACKAGE_DIR)
	cp ./CONTRIBUTING.md $(PACKAGE_DIR)
	cp ./LICENSE $(PACKAGE_DIR)
	cp ./cmd/trickster/conf/*.conf $(CONF_DIR)
	
	GOOS=darwin  GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(BIN_DIR)/trickster-$(PROGVER).darwin-amd64  -a -v $(TRICKSTER_MAIN)/*.go
	GOOS=linux   GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(BIN_DIR)/trickster-$(PROGVER).linux-amd64   -a -v $(TRICKSTER_MAIN)/*.go
	GOOS=linux   GOARCH=arm64 $(GO) build $(LDFLAGS) -o $(BIN_DIR)/trickster-$(PROGVER).linux-arm64   -a -v $(TRICKSTER_MAIN)/*.go
	GOOS=windows GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(BIN_DIR)/trickster-$(PROGVER).windows-amd64 -a -v $(TRICKSTER_MAIN)/*.go

	cd ./$(BUILD_SUBDIR) && tar cvfz ./trickster-$(PROGVER).tar.gz ./trickster-$(PROGVER)/*

# Minikube and helm bootstrapping are done via deploy/helm/Makefile
.PHONY: helm-local
helm-local:
	kubectl config use-context minikube --namespace=trickster
	kubectl scale --replicas=0 deployment/dev-trickster -n trickster
	eval $$(minikube docker-env) \
		&& docker build -f deploy/Dockerfile -t trickster:dev .
	kubectl set image deployment/dev-trickster trickster=trickster:dev -n trickster
	kubectl scale --replicas=1 deployment/dev-trickster -n trickster

# Minikube and helm bootstrapping are done via deploy/kube/Makefile
.PHONY: kube-local
kube-local:
	kubectl config use-context minikube
	kubectl scale --replicas=0 deployment/trickster
	eval $$(minikube docker-env) \
		&& docker build -f deploy/Dockerfile -t trickster:dev .
	kubectl set image deployment/trickster trickster=trickster:dev
	kubectl scale --replicas=1 deployment/trickster

.PHONY: docker
docker:
	docker build --build-arg IMAGE_ARCH=$(IMAGE_ARCH) --build-arg GOARCH=$(GOARCH) -f ./deploy/Dockerfile -t trickster:$(PROGVER) .

.PHONY: docker-release
docker-release:
# linux x86 image
	docker build --build-arg IMAGE_ARCH=amd64 --build-arg GOARCH=amd64 -f ./deploy/Dockerfile -t tricksterio/trickster:$(IMAGE_TAG) .
# linux arm image
	docker build --build-arg IMAGE_ARCH=arm64v8 --build-arg GOARCH=arm64 -f ./deploy/Dockerfile -t tricksterio/trickster:arm64v8-$(IMAGE_TAG) .

.PHONY: style
style:
	! gofmt -d $$(find . -path ./vendor -prune -o -name '*.go' -print) | grep '^'

.PHONY: test
test:
	@go test -v -coverprofile=.coverprofile ./... 

.PHONY: bench
bench:
	bash -c "$(GO) test -v -coverprofile=.coverprofile ./... -run=nonthingplease -bench=. | grep -v ' app=trickster '; exit ${PIPESTATUS[0]}"

.PHONY: test-cover
test-cover: test
	$(GO) tool cover -html=.coverprofile

.PHONY: clean
clean:
	rm -rf ./trickster ./$(BUILD_SUBDIR)
