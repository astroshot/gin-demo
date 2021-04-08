package util

const (
	// SuccessInfo defines successful HTTP response hint
	SuccessInfo = "Success"

	// FailInfo defines failed HTTP response hint
	FailInfo = "Failure"

	// URINotFound defines hint message when HTTP URI not found
	URINotFound = "Method not found"
	// MethodNotSupported defines hint message when HTTP method not implemented
	MethodNotSupported = "Method not supported"

	// DefaultDateTimeFormat defines default datetime format
	DefaultDateTimeFormat          = "2006-01-02 15:04:05"
	DateTimeFormatWithMicroseconds = "2006-01-02 15:04:05.000000"

	TraceIDKey = "gin-demo/traceID"
	EmptyStr   = ""

	TraceIDField = "TraceID"
)
