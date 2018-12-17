
.ONESHELL:
build:*
	GOOS=linux GOARCH=amd64 go build main.go 
	docker-compose build
run:
	docker-compose up -d
