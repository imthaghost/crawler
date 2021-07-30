.PHONY: binary clean package clean-package test

BINARY_NAME = crawler
BUILD_DIR = ./bin

ifeq ($(BUILD_ID),)
DOCKERFILE := dev.dockerfile
else
DOCKERFILE := ci.dockerfile
endif

ifeq ($(TAG),)
TAG := $(shell openssl md5 $(DOCKERFILE) | cut -f 2 -d " ")
endif

DOCKER_IMG = imthaghost/crawler:$(TAG)

DOCKER_SH = docker run \
		-v "$(shell pwd):/workdir:rw" \
		-w /workdir \
    	-t "$(DOCKER_IMG)" \
    	sh -c

# DEFAULT GOAL
binary: $(BUILD_DIR)/crawler

# make the tester image
.tester_image:
	docker build -t "$(DOCKER_IMG)" . -f $(DOCKERFILE)
	touch $@

# lint code after building the image
.lint: .tester_image
	$(DOCKER_SH) "go vet ./..."
	#$(DOCKER_SH) "golangci-lint run -v"
	touch $@

.test: .lint
	$(DOCKER_SH) "go test -v ./..."
	touch $@

# build the binary after testing
$(BUILD_DIR)/crawler: .test
	$(DOCKER_SH) "GOOS=linux go build -o ./bin/crawler -mod vendor \
	github.com/imthaghost/crawler/cmd/crawler"

package:
	DOCKER_IMG=$(DOCKER_IMG) make -C package/

clean:
	rm -vfr ./bin/crawler
	rm -vfr .tester_image
	rm -vfr .lint
	rm -vfr .test
