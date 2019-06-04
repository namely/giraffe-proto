FROM namely/protoc-all:1.21_0

COPY . .

RUN entrypoint.sh -f namely/giraffe/stitch.proto -l go
