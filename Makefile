.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/check-schedule check/*

test:
	go test ./...

clean:
	rm -rf ./bin

deploy: clean build test
	sls deploy --verbose
