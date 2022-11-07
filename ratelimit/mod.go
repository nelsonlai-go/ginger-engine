package ratelimit

import (
	"log"
	"time"

	"github.com/nelsonlai-go/ginger-engine/ginger"
)

type ConfigOption struct {
	GeneralRateLimit         int64
	GeneralRateLimitDuration time.Duration
}

var option *ConfigOption = nil

func Config(opt *ConfigOption) {
	option = opt
}

func Register(e ginger.Ginger) {
	if option == nil {
		log.Fatalln("ratelimit: Option is nil")
	}
	e.Middleware(RateLimit(option.GeneralRateLimitDuration, option.GeneralRateLimit))
}
