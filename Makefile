build:
	go build -o bin/portfolio cmd/portfolio/main.go
	du -h bin/portfolio

build-prod:
	go build -ldflags "-w" -o bin/portfolio cmd/portfolio/main.go
	du -h bin/portfolio

start:
	bin/portfolio

dev:
	go run cmd/portfolio/main.go
