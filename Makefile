format:
	cd clientapi && go fmt

lint:
	golint ./clientapi

vet:
	cd clientapi && go vet

unit-test:
	cd clientapi && go test -v ./...

all-locally: format lint vet unit-test
