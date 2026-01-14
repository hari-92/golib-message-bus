package properties

import (
	"github.com/golibs-starter/golib/config"
	"github.com/hari-92/golib-message-bus/kafka/core"
)

func NewTopicAdmin(loader config.Loader) (*TopicAdmin, error) {
	props := TopicAdmin{}
	err := loader.Bind(&props)
	return &props, err
}

type TopicAdmin struct {
	Topics []core.TopicConfiguration
}

func (h TopicAdmin) Prefix() string {
	return "app.kafka.admin"
}
