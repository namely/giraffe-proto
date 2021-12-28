ARG PROTOC_VERSION = 1.42_0

FROM namely/protoc-all:${PROTOC_VERSION}

COPY . .

RUN entrypoint.sh -f namely/giraffe/stitch.proto -l go
