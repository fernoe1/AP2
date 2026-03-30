package http

import (
	"net/http"
	"time"
)

var C *http.Client

func init() {
	C = &http.Client{Timeout: 2 * time.Second}
}
