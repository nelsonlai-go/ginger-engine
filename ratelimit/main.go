package ratelimit

import (
	"time"

	"github.com/nelsonlai-go/ginger-engine/ginger"
)

const (
	GENERAL_RATE_LIMIT_PERIOD = "GENERAL_RATE_LIMIT_PERIOD"
	GENERAL_RATE_LIMIT_LIMIT  = "GENERAL_RATE_LIMIT_LIMIT"
)

func RegisterHandler(e ginger.Ginger, option ginger.RegisterOption) {
	e.Middleware(RateLimit(
		option.Param(GENERAL_RATE_LIMIT_PERIOD, true).(time.Duration),
		option.Param(GENERAL_RATE_LIMIT_LIMIT, true).(int64),
	))
}
