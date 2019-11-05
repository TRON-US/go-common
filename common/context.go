package common

// gracefully deprecating this package
// moving forward to use constant package instead

import (
	"github.com/tron-us/go-common/constant"
)

const (
	ContextHandlerKey = constant.HandlerNameContext
	ContextHTTPURLKey = constant.HTTPURLContext
)
