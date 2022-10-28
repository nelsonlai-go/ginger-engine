package cors

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/nelsonlai-go/ginger-engine/ginger"
)

const (
	ALLOW_ALL_ORIGINS        = "ALLOW_ALL_ORIGINS"
	ALLOW_ORIGINS            = "ALLOW_ORIGINS"
	ALLOW_ORIGIN_FUNC        = "ALLOW_ORIGIN_FUNC"
	ALLOW_METHODS            = "ALLOW_METHODS"
	ALLOW_HEADERS            = "ALLOW_HEADERS"
	ALLOW_CREDENTIALS        = "ALLOW_CREDENTIALS"
	EXPOSE_HEADERS           = "EXPOSE_HEADERS"
	MAX_AGE                  = "MAX_AGE"
	ALLOW_WILD_CARD          = "ALLOW_WILD_CARD"
	ALLOW_BROWSER_EXTENSIONS = "ALLOW_BROWSER_EXTENSIONS"
	ALLOW_WEB_SOCKETS        = "ALLOW_WEB_SOCKETS"
	ALLOW_FILES              = "ALLOW_FILES"
)

func RegisterHandler(e ginger.Ginger, option ginger.RegisterOption) {

	corsOpt := cors.Config{}

	if option.Param(ALLOW_ALL_ORIGINS, false) != nil {
		corsOpt.AllowAllOrigins = option.Param(ALLOW_ALL_ORIGINS, false).(bool)
	}

	if option.Param(ALLOW_ORIGINS, false) != nil {
		corsOpt.AllowOrigins = option.Param(ALLOW_ORIGINS, false).([]string)
	}

	if option.Param(ALLOW_ORIGIN_FUNC, false) != nil {
		corsOpt.AllowOriginFunc = option.Param(ALLOW_ORIGIN_FUNC, false).(func(origin string) bool)
	}

	if option.Param(ALLOW_METHODS, false) != nil {
		corsOpt.AllowMethods = option.Param(ALLOW_METHODS, false).([]string)
	}

	if option.Param(ALLOW_HEADERS, false) != nil {
		corsOpt.AllowHeaders = option.Param(ALLOW_HEADERS, false).([]string)
	}

	if option.Param(ALLOW_CREDENTIALS, false) != nil {
		corsOpt.AllowCredentials = option.Param(ALLOW_CREDENTIALS, false).(bool)
	}

	if option.Param(EXPOSE_HEADERS, false) != nil {
		corsOpt.ExposeHeaders = option.Param(EXPOSE_HEADERS, false).([]string)
	}

	if option.Param(MAX_AGE, false) != nil {
		corsOpt.MaxAge = option.Param(MAX_AGE, false).(time.Duration)
	}

	if option.Param(ALLOW_WILD_CARD, false) != nil {
		corsOpt.AllowWildcard = option.Param(ALLOW_WILD_CARD, false).(bool)
	}

	if option.Param(ALLOW_BROWSER_EXTENSIONS, false) != nil {
		corsOpt.AllowBrowserExtensions = option.Param(ALLOW_BROWSER_EXTENSIONS, false).(bool)
	}

	if option.Param(ALLOW_WEB_SOCKETS, false) != nil {
		corsOpt.AllowWebSockets = option.Param(ALLOW_WEB_SOCKETS, false).(bool)
	}

	if option.Param(ALLOW_FILES, false) != nil {
		corsOpt.AllowFiles = option.Param(ALLOW_FILES, false).(bool)
	}

	e.Middleware(cors.New(corsOpt))
}
