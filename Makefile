run:
	go run src/cmd/api/main.go

build:
	go build -o src/cmd/api/rest-api src/cmd/api/main.go

compose-up:
	docker-compose -f deployment/docker-compose.yml up -d --build

compose-up-app-only:
	docker-compose -f deployment/docker-compose.yml up -d --build app

compose-up-no-app:
	docker-compose -f deployment/docker-compose.yml up -d postgres

compose-down:
	docker-compose -f deployment/docker-compose.yml down
