package helpers

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

func GetRemoteAddress(r *http.Request) string {
	if val := r.Header.Get("X-Forwarded-For"); val != "" {
		return strings.Split(val, ":")[0]
	} else if val := r.Header.Get("X-Real-IP"); val != "" {
		return strings.Split(val, ":")[0]
	} else {
		return strings.Split(r.RemoteAddr, ":")[0]
	}
}

func VerifyMinLen(m map[string]string) (names string, err error) {

	for k, v := range m {
		if len(v) == 0 {
			err = errors.New("len == 0")
			names += k
			names += " "
		}
	}

	return
}

func GetTimeNowString() string {
	return time.Now().Format("2006-01-02 15:04:05.000000")
}

func ConvertStringToUUID(str string) (id uuid.UUID, err error) {
	id, err = uuid.FromString(str)
	return
}
