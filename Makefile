
run:
	go run ./cmd/neocal.go

test:
	go test ./...

.PHONY: build
build:
	rm -rf ./build/
	mkdir -p ./build/
	go build \
		-o ./build/neocal \
		-gcflags -m=2 \
		./cmd/ 
