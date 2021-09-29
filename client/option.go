package client

import (
	"fmt"
	"time"
)

const (
	clientVersion = "v0.0.1"

	defaultBaseURL = "https://minfraud.maxmind.com/minfraud"
	defaultTimeout = 20 * time.Second
)

var defaultUserAgent = fmt.Sprintf("minfraund-api-go/%s", clientVersion)

// Option contains optional setting of Client.
type Option struct {
	BaseURL   string
	UserAgent string
	Timeout   time.Duration
	Debug     bool
	Retry     bool
	LogFn     func(msg string, opts ...interface{})
}

func (o Option) LogInfo(msg string, opts ...interface{}) {
	if o.LogFn == nil {
		return
	}
	o.LogFn(msg, opts...)
}

func (o Option) getBaseURL() string {
	if o.BaseURL != "" {
		return o.BaseURL
	}
	return defaultBaseURL
}

func (o Option) getUserAgent() string {
	if o.UserAgent != "" {
		return o.UserAgent
	}
	return defaultUserAgent
}

func (o Option) getTimeout() time.Duration {
	if o.Timeout > 0 {
		return o.Timeout
	}
	return defaultTimeout
}
