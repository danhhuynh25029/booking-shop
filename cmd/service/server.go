package main

import (
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"services/pkg/constant"
	"services/pkg/kafka"
)

func main() {
	r := gin.Default()
	producer, err := kafka.NewProducer()
	if err != nil {
		log.Printf("function NewProducer err : %v", err)
	}
	consumer, err := sarama.NewConsumer(constant.Brokers, nil)
	if err != nil {
		log.Printf("function NewConsumer err : %v", err)
	}
	kafka.Subscribe(constant.Topic, consumer)

	r.GET("/ping", func(context *gin.Context) {
		message := kafka.PrepareMessage(constant.Topic, "hello")
		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Printf("%s error occured.", err.Error())
		} else {
			log.Printf("Message was saved to partion: %d.\nMessage offset is: %d.\n", partition, offset)
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err = r.Run()
	if err != nil {
		log.Fatalf("cannot start serve : %v", err)
	}
}
