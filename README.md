# services

# Kafka
* Start kafka with docker [Document](https://developer.confluent.io/quickstart/kafka-docker/)
* Create topic in kafka

``
  docker exec broker \
  kafka-topics --bootstrap-server broker:9092 \
  --create \
  --topic quickstart
``