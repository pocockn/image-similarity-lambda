
BRANCH = "master"
PACKAGES ?= "./..."
VERSION = $(shell cat ./VERSION)

DEFAULT: test

build-binary:
	@echo "=> Building binary"
	@GOOS=linux go build -o image-similarity-analysis -i -ldflags "-X main.Version=${VERSION}"

check-version:
	git fetch && (! git rev-list ${VERSION})

push-tag:
	git checkout ${BRANCH}
	git pull origin ${BRANCH}
	git tag ${VERSION}
	git push origin ${BRANCH} --tags

test:
	@go test "${PACKAGES}" -cover

vet:
	@go vet "${PACKAGES}"

zip-binary: build-binary
	@echo "Zipping binary..."
	zip -r image-similarity-analysis image-similarity-analysis