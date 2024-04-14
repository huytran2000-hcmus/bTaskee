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

## audit: clean up all your code in
.PHONY: audit
audit: test
	@echo 'Tidying and verifying module dependencies...'
	cd booking; go mod tidy; go mod verify
	cd pricing; go mod tidy; go mod verify
	cd send; go mod tidy; go mod verify

## test: run all test
.PHONY: test
test:
	@echo 'Runnings tests'
	cd booking; go test -vet=off --count=1 ./...
	cd pricing; go test -vet=off --count=1 ./...
	cd send; go test -vet=off  --count=1 ./...


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

## mongosh dbname=$1: spin down mongodb container
.PHONY: mongosh
mongosh:
	docker exec -it dev-mongo mongosh -u btaskee -p secret --authenticationDatabase admin ${dbname}

## build-images: build all docker images
.PHONY: build-images
build-images:
	cd booking; docker build -t huypk2000/booking:latest .
	cd pricing; docker build -t huypk2000/pricing:latest .
	cd send; docker build -t huypk2000/send:latest .

## swag: generate all swagger docs
.PHONY: swag
swag:
	cd booking; swag init -g ./cmd/main.go
	cd pricing; swag init -g ./cmd/main.go
	cd send; swag init -g ./cmd/main.go

## gen-mock: generate mocks
.PHONY: gen-mock
gen-mock:
	cd booking; mockery
	cd pricing; mockery
	cd send; mockery

## goconvey: run goconvey
.PHONY: goconvey
goconvey:
	goconvey --port 8800
