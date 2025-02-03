
test:
	go test -v ./... -run=.*/february-with-28-days_eu_00
	@#go test -v ./... -run=.*/february-with-28-days_us_00
	@#go test -v ./...
	@#go test ./...

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
