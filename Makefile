
run:
	go run ./cmd/cal.go

test:
	go test ./...

.PHONY: build
build:
	rm -rf ./build/
	mkdir -p ./build/
	go build \
		-o ./build/cal \
		-gcflags -m=2 \
		./cmd/ 
