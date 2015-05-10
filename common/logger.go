package common

//	log.Debug("debug %s", "hoge")
//	log.Info("info")
//	log.Notice("notice")
//	log.Warning("warning")
//	log.Error("error")
//	log.Critical("critical")

import (
	"github.com/op/go-logging"
)

var instance *logging.Logger

//var logger = logging.MustGetLogger("example")
var format = logging.MustStringFormatter(
	"%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}",
)

func initLogger() {
	logging.SetFormatter(format)
	logging.SetLevel(logging.DEBUG, "")
}

func newLogger() *logging.Logger {
	logger := logging.MustGetLogger("example")
	initLogger()

	return logger
}

func GetLogger() *logging.Logger {
	if instance == nil {
		instance = newLogger()
	}
	return instance
}
