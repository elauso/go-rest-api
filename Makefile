docker-build:
	docker-compose up --build

docker-run:
	docker-compose up

docker-clean:
	docker-compose down -v

build:
	go build -o ./bin/go-rest-api ./src/net/elau/gorestapi/server.go

run:
	./bin/go-rest-api

clean:
	rm -rf ./bin