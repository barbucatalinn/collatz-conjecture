.PHONY: docker test check staticcheck golint

docker:
	docker-compose run --service-ports app bash

test:
	gotest -v -cover ./...

golint:
	golint ./...

govet:
	go vet ./...

staticcheck:
	staticcheck ./...

check: test golint govet staticcheck

run:
	go run cmd/app/main.go