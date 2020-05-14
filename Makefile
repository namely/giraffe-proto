GRPC_VERSION = 1.29_1

.PHONY: protos
protos:
	docker run --rm -v ${PWD}:/defs namely/protoc-all:${GRPC_VERSION} -f namely/giraffe/stitch.proto -l go -o .
	mv namely/giraffe/*.go ${PWD}
