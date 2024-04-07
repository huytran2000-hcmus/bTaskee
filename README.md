# bTaskee
An on-demand home cleaning service

## How does it work?
The booking service will call the pricing service to calculate the price of a task and save the task in the database
(MongoDB). Then it will call send service to send tasker email.


## How to run it?
There are 3 ways:
- Use docker-compose
- Manually run 3 services and Mongodb in 4 separate terminals
```bash
	cd booking; go run cmd/main.go
```

```bash
	cd pricing; go run cmd/main.go
```

```bash
	cd send; go run cmd/main.go
```

```bash
	docker run \
		-d \
		--name dev-mongo \
		-p 27017:27017 \
		-e MONGO_INITDB_ROOT_USERNAME=btaskee \
		-e MONGO_INITDB_ROOT_PASSWORD=secret \
		mongo:7.0.8
```
All of these commands have their corresponding target in Makefile (use make help to get more info)
## How to use it?
Access to swagger links:
[](http://localhost:8080/swagger/index.html)
[](http://localhost:8081/swagger/index.html)
[](http://localhost:8082/swagger/index.html)

There are ready-to-use example requests for you to test.
