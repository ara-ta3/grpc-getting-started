GO=go
PROTO=protoc

.PHONY: proto

run:
	$(GO) run ./main.go

proto:
	$(PROTO) --go_out=plugins=gprc:./proto *.proto
