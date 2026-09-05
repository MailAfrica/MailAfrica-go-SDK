package mailafrica

import (
	"net/http"
	"time"
)

// Hooks provides optional observability hooks.
type Hooks struct {
	OnRequest  func(req *http.Request)
	OnResponse func(resp *http.Response, duration time.Duration)
	OnError    func(err error)
}
