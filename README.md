# Rabbit MQ Go

Experimenting messaging queue with RabbitMQ in Go

1. Set the enviroment variables
```shell
make env
```

2. Install husky to prevent bad commit messages
```shell
make ih
```

3. build the RabbitMQ container
```shell
make rmq
```

## Useful RabbitMQ Commands
1. List all queues:
```shell
rabbitmqctl list_queues
```

2. Delete a queue:
```shell
rabbitmqctl delete_queue <queue-name>
```

3. Get all the messages sended but not yet akcnowledged by the consumers:
```shell
rabbitmqctl list_queues <queue-name> messages_ready messages_unacknowledged
```
