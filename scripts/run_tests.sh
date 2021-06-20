#!/bin/sh

golint ./clientapi
cd clientapi && go vet
cd ..
cd clientapi && go test -v ./...
cd ..
