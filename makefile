run:
	go run ./cmd/web/

live:
	nodemon --exec go run ./cmd/web/. --ext go

up:
	 cd ./sql/schema && goose postgres postgres://root:root@localhost:5432/go-ws up

down:
	 cd ./sql/schema && goose postgres postgres://root:root@localhost:5432/go-ws down

build:
	go build -C cmd/web -o ../../go-ws.exe

build-run:
	go build -C cmd/web -o ../../go-ws.exe && ./go-ws.exe