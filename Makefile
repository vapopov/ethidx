export GO111MODULE=on
export CGO_ENABLED=0

.PHONY token:
token:
	abigen --abi abi/contract.abi --pkg ethidx --type Token --out token.go

vendor:
	go mod vendor

.PHONY run:
run:
	go run ./cmd/ethidx/

.PHONY build:
build:
	go build -o ethidx ./cmd/ethidx

.PHONY swag:
swag:
	swag init -g api.go


.PHONY docker:
docker: token swag
	docker build . -t zion/ethidx:latest
