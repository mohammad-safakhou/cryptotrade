package pkg

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

type MessageBrokerHandler interface {
	Push(ctx context.Context, data []byte) error
	Consumer(ctx context.Context, dataChan chan []byte)
}

type kafkaHandler struct {
	kafkaConn *kafka.Conn
}

func NewKafkaHandler(kafkaConn *kafka.Conn) MessageBrokerHandler {
	return &kafkaHandler{kafkaConn: kafkaConn}
}

func (kf *kafkaHandler) Push(ctx context.Context, data []byte) error {
	_, err := kf.kafkaConn.WriteMessages(kafka.Message{Value: data})
	if err != nil {
		return err
	}
	return nil
}

func (kf *kafkaHandler) Consumer(ctx context.Context, dataChan chan []byte) {
	batch := kf.kafkaConn.ReadBatch(10e3, 1e6)

	b := make([]byte, 10e3)
	for {
		n, err := batch.Read(b)
		if err != nil {
			fmt.Println(err)
		}
		dataChan <- b[:n]
	}
}

func KafkaConnection(host, port, topic string, partition int) (*kafka.Conn, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", host+":"+port, topic, partition)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
