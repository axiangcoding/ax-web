run:
	pwd&&cd cmd/app/ && go run .
.PHONY: run

.PHONY: generate

generate:
	cd cmd/app&&wire

.PHONY: proto
proto:
	cd api/app/pb&&protoc --go_out=plugins=grpc:. *.proto