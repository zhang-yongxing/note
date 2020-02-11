package utils

import (
	uuid "github.com/satori/go.uuid"
	"strings"
)

func UStr32()  string {
	us := uuid.Must(uuid.NewV4()).String()
	return strings.ReplaceAll(us, "-", "")
}