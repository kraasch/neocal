
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

hub_update:
	@hub_ctrl ${HUB_MODE} ln "$(realpath ./build/neocal)"
