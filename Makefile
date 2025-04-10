# Go parameters
GOCMD:=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

server-run:
	$(GORUN) server/main.go

publish-run:
	$(GORUN) publish/main.go

subscribe-run:
	$(GORUN) subscribe/main.go $(ARGS)

test:
	$(GOCMD) clean -testcache
	$(GOTEST) ./...

testc:
	$(GOCMD) clean -testcache
	$(GOTEST) -cover ./...

protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protobuf/**/*.proto
