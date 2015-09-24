package setting

import (
	"github.com/Unknwon/log"
	"os"
)

var (
	HTTPPort = os.Getenv("PORT")
)

func init() {
	log.Prefix = "[wangqi]"
}
