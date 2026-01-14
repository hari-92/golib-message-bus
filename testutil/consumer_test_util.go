package golibmsgTestUtil

import (
	"fmt"
	"strings"
	"time"

	"github.com/hari-92/golib-message-bus/kafka/core"
)

// HandleMessage handle message
func HandleMessage(consumerName string, message []byte) {
	consumer, ok := consumerMap[strings.ToLower(consumerName)]
	if !ok {
		panic(fmt.Sprintf("consumer with name %v not found", consumerName))
	}
	consumer.HandlerFunc(&core.ConsumerMessage{
		Topic:     "kafka-consumer-test-util-topic",
		Key:       nil,
		Value:     message,
		Headers:   nil,
		Partition: 0,
		Offset:    0,
		Timestamp: time.Now(),
	})
}
