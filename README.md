# Foroom Project

## Description
Foorom project is a dummy project capabilities of serving a forum between multiple users

## Dependencies
- Go
- Kafka (sarama)

## Starting the app
1. Run your Zookeeper server
    ```sh
    bin/zookeeper-server-start.sh config/zookeeper.properties
    ```

2. Run your Kafka server
    ```sh
    bin/kafka-server-start.sh config/server.properties
    ```

3. Create a topic (room name)
    ```sh
    bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic <topic_name>
    ```

4. Run Foroom App
    ```sh
    go run main.go -user=<username> -room=<room_name/topic>
    ```

## Author
Ferdinandus Richard
