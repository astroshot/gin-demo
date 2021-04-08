package helper

import (
	_ "math/rand"
	_ "strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func MakeUUID() string {
	uid := uuid.NewV4()
	return uid.String()
}

func MakeTraceID() string {
	var uuid = MakeUUID()
	return strings.Replace(uuid, "-", "", -1)
}
