PROJECTNAME=$(shell basename "$(PWD)")
GOBASE=$(PWD)
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)

RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(RUN_ARGS):;@:)

createdb:
	docker run -d --name mydb -p 5123:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=test postgres:alpine
	sleep 5

## start: Run the application
start:
	DATABASE_URL="postgres://admin:password@127.0.0.1:5432/test?sslmode=disable" go run main.go

run: start


## tests: Run all tests
tests:
	test_database_url="postgres://admin:password@127.0.0.1:5432/test?sslmode=disable" go test -v --cover ./src/app/...

test: tests

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
