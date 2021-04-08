package helper

import (
	"context"

	"gin-demo/pkg/util"
)

func GetTraceIDFrom(ctx context.Context) string {
	var traceVal = ctx.Value(util.TraceIDKey)
	traceID, ok := traceVal.(string)
	if !ok {
		return util.EmptyStr
	}

	return traceID
}
