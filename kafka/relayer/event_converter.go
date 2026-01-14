package relayer

import (
	"github.com/golibs-starter/golib/pubsub"
	"github.com/hari-92/golib-message-bus/kafka/core"
)

type EventConverter interface {

	// Convert internal Event to kafka message
	Convert(event pubsub.Event) (*core.Message, error)

	// Restore a consumed message back to destination event
	Restore(msg *core.ConsumerMessage, dest pubsub.Event) error
}
