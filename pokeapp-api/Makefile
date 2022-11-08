BINARY=pokeapp-api 

VERSION=v1.0.0

# introduce informacion en el binario
# LDFLAGS=-ldflags="-X 'main.Version=${VERSION}'"

build:
	go build -o ${BINARY} ./main.go

run:
	go run main.go

up:
	docker-compose up --build 
 
down:
	docker-compose down --remove-orphans
	

migrate_down:
	migrate -database "postgres://postgres:root@localhost:1234/pokeapp?sslmode=disable" -path migrations down 1

migrate_up:
	migrate -database "postgres://postgres:root@localhost:1234/pokeapp?sslmode=disable" -path migrations up 1

