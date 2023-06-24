GOPATH=$(shell go env GOPATH)

test:
	go test -covermode=atomic -cover -v ./... -coverprofile=coverage.out

vet:
	go vet ./...

staticcheck:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	$(GOPATH)/bin/staticcheck ./...

fmt:
	go fmt ./...

tidy:
	go mod tidy

polish: tidy fmt vet staticcheck

.PHONY: polish tidy fmt vet staticcheck test
