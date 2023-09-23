build:
	go build -o bin/portfolio cmd/portfolio/main.go
	du -h bin/portfolio

start:
	bin/portfolio

dev:
	go run cmd/portfolio/main.go
