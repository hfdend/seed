package cli

import (
	"seed/conf"
	"seed/logger"
)

var (
	Logger *logger.Logger
)

func InitializeLogger() {
	dir := "/dev/stdout"
	if conf.Config.Log.Dir != "" {
		dir = conf.Config.Log.Dir
	}
	Logger = logger.NewLogger(dir, "20060102").Kind("seed")
}
