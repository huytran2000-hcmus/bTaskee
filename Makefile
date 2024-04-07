## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## run/booking: run the booking application
.PHONY: run/booking
run/booking:
	cd booking; go run cmd/main.go
	
## run/pricing: run the pricing application
.PHONY: run/pricing
run/pricing:
	cd pricing; go run cmd/main.go

## run/send: run the send application
.PHONY: run/send
run/send:
	cd send; go run cmd/main.go

## audit: clean up all your code
.PHONY: audit
audit: test document
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify

## test: run all test
.PHONY: test
test:
	@echo 'Runnings tests'
	go test -vet=off ./...

## document: generate the swagger documentation
.PHONY: document
document:
	@echo "Generate swagger documentation..."
	swag init

## mongodb\up: spin up mongodb container
.PHONY: mongodb/up
mongodb/up:
	@echo "Spin up mongodb container"
	docker run \
		-d \
		--name dev-mongo \
		-p 27017:27017 \
		-e MONGO_INITDB_ROOT_USERNAME=btaskee \
		-e MONGO_INITDB_ROOT_PASSWORD=secret \
		mongo:7.0.8

## mongodb\down: spin down mongodb container
.PHONY: mongodb/down
mongodb/down:
	docker stop dev-mongo
	docker rm dev-mongo

## mongosh: spin down mongodb container
.PHONY: mongosh
mongosh:
	docker exec -it dev-mongo mongosh -u btaskee -p secret --authenticationDatabase admin ${dbname}
