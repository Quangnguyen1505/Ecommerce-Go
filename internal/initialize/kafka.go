package initialize

import (
	"log"

	"github.com/ntquang/ecommerce/global"
	kafka "github.com/segmentio/kafka-go"
)

// init kafka producer
func InitKafa() {
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP("localhost:19092"),
		Topic:    "otp-auth-topic",    //topic nhan otp
		Balancer: &kafka.LeastBytes{}, // can bang tai
	}
}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		log.Fatalf("Failed to close kafka producer: %v", err)
	}
}
