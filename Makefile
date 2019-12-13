default:
	go test -cpu=1,2,4 -v -tags integration ./...

vet:
	go vet $(go list ./... | grep -v /vendor/)

tidy:
	go mod tidy

vendor:
	go mod vendor

.PHONY: default vet tidy vendor
