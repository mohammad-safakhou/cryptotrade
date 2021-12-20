package pkg

import (
	"context"
	"github.com/segmentio/kafka-go"
)

func KafkaConnection(host, port, topic string, partition int) (*kafka.Conn, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
