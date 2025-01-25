package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

var (
	kafkaURL   = "localhost:19092"
	kafkaTopic = "user_topic_vip"
)

// for producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{}, // can bang tai
	}
}

// for consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers, // danh sach brokers example: []string{"localhost:1", "localhost:2"}
		GroupID:        groupID, // vung consume
		Topic:          topic,
		MinBytes:       10e3,             // 10KB van chuyen du lieu
		MaxBytes:       10e6,             // 10MB
		CommitInterval: time.Second,      // khoang thoi gian giua cac commit 1s
		StartOffset:    kafka.LastOffset, //offset cuoi k lay msg cu //kafka.FirstOffset, // dat gia tri ban dau khi consumer bat dau lang nghe
	})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// mua ban chung khoan
func newStock(msg, typeMsg string) *StockInfo {
	s := StockInfo{}
	s.Message = msg
	s.Type = typeMsg

	return &s
}

func actionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))
	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s

	jsonBody, _ := json.Marshal(body)

	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}
	//write message by producer
	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
		"msg": "action successfully",
	})
}

// Consumer listen buy ATC
func RegisterConsumerATC(id int) {
	// group consumer
	kafkaGroupId := fmt.Sprintf("consumer-group-%d", id)
	// dk đọc ở topic, groupid nào
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	defer reader.Close() // giai phong ket noi sau khi ham chạy thanh cong hay that bai

	fmt.Println("Consumer (", id, ") Listen Session ATC")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer(%d) error: %v", id, err)
		}
		fmt.Printf("Consumer (%d), listen topic: %v, partition: %v, offset: %v, time: %d %s = %s\n", id, m.Topic, m.Partition,
			m.Offset, m.Time.Unix(), string(m.Key), string(m.Value))
	}
}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()

	r.POST("action/stock", actionStock)

	// register listen
	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)
	go RegisterConsumerATC(3)
	go RegisterConsumerATC(4)
	r.Run(":8989")
}
