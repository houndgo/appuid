package appuid

import (
	"strings"

	"github.com/satori/go.uuid"
)

// CreateAPPUUID is
func CreateAPPUUID() string {
	u := uuid.Must(uuid.NewV4())
	us := strings.Replace(u.String(), "-", "", -1)
	return us
}
