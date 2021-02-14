package helper

import (
	_ "math/rand"
	_ "strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func GetUUID() string {
	uid := uuid.NewV4()
	return uid.String()
}

func GetTraceID() string {
	var uuid = GetUUID()
	return strings.Replace(uuid, "-", "", -1)
}
