run:
	go run ./cmd/web/

live:
	nodemon --exec go run ./cmd/web/. --ext go