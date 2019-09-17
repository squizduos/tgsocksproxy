
deps:
	go get -d -v ./...

build:
	go build -ldflags "-linkmode external -extldflags -static" -o ./bin/proxy

run: build
	DEBUG=true SOCKS_HOST=bot.localtest.me ./bin/proxy
