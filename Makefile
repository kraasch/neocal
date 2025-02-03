
test:
	@#go test -v ./...
	@#go test -v ./pkg/calengine/...
	go test -v ./pkg/calcontroller/...

run:
	go run ./cmd/cal.go

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
