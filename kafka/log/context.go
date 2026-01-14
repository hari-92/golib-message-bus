package log

import (
	"github.com/golibs-starter/golib/log/field"
	"github.com/golibs-starter/golib/web/constant"
	webLog "github.com/golibs-starter/golib/web/log"
	kafkaConstant "github.com/hari-92/golib-message-bus/kafka/constant"
)

func GetLoggingContext(metadata map[string]interface{}) []field.Field {
	loggingCtx := metadata[kafkaConstant.LoggingContext]
	if loggingCtx == nil {
		return nil
	}
	ctx, ok := loggingCtx.(*webLog.ContextAttributes)
	if !ok {
		return []field.Field{}
	}
	return []field.Field{field.Object(constant.ContextReqMeta, ctx)}
}
