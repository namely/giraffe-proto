PROTOC_VERSION = 1.42_0

.PHONY: protos
protos:
	docker run --rm -v ${PWD}:/defs namely/protoc-all:${PROTOC_VERSION} -f namely/giraffe/stitch.proto -l go -o .
	mv github.com/namely/giraffe-proto/*.go ${PWD}
	rmdir github.com/namely/giraffe-proto github.com/namely github.com
