docker:
	docker-compose run -p 6060:6060 --rm my-go-project /bin/bash

run:
	go run src/my-lib/main.go

docs:
	godoc -src -http ":6060"
