package cors

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/nelsonlai-go/ginger-engine/ginger"
)

var option *cors.Config

func Config(config *cors.Config) {
	option = config
}

func Register(e ginger.Ginger) {
	if option == nil {
		log.Fatalln("cors: option is nil")
	}
	e.Middleware(cors.New(*option))
}
