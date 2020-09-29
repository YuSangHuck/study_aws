.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/main com/handler.go
	npm install

clean:
	rm -rf ./bin ./vendor Gopkg.lock
	rm -rf node_modules

deploy: clean build
	serverless deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
