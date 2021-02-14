package util

const (
	// SuccessInfo defines successful HTTP response hint
	SuccessInfo = "Success"

	// FailInfo defines failed HTTP response hint
	FailInfo = "Fail"

	// MethodNotFound defines hint message when HTTP method not implemented
	MethodNotFound = "Method not found"

	// DefaultDateTimeFormat defines default datetime format
	DefaultDateTimeFormat          = "2006-01-02 15:04:05"
	DateTimeFormatWithMicroseconds = "2006-01-02 15:04:05.000000"

	TraceIDKey = "gin-demo/traceID"
)
