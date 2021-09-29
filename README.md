# Connecting to RabbitMQ with GO

In order to run this application, you'll need to first run RabbitMQ in a docker container:

```shell
docker run -d --hostname rabbitmq --name test-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```

Once docker is up, you can ensure RMQ is working by visiting http://localhost:15672/ in your browser. You should get the RMQ admin panel. 

The login for the admin panel is 'guest', password is also 'guest'.


After you've ensured RMQ is working, you can run the app with:  
`go run ./cmd/server/main.go`

You should see the message 'Hi' printed to your console.

In the [rabbitmq.go](https://github.com/millbj92/go-rabbitmq/blob/a137c4971781b6b5c64b616a85a062ff490e9c5b/internal/rabbitmq/rabbitmq.go#L50-L88) file you can see that we are starting as both a publisher and a subscriber to the same queue. 

If you'd like to publish a message to the queue to watch it be consumed, run `docker ps` to grab your docker container name and then run the following:

```shell
docker exec -it DOCKER_CONTAINER_NAME bash   
```

once logged into your container, you can run:

```shell
rabbitmqadmin publish exchange=amq.default routing_key="TestQueue" payload="Your Message Here"
```

This was a pretty cool exercise. I think I'll do Kafka next ;)
