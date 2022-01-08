package logging

import log "github.com/sirupsen/logrus"

func SetupLogger() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "01-02 15:04:05.000"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
	log.SetLevel(log.DebugLevel)
}
