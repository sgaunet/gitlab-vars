package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func initTrace(outputFile *os.File, debugLevel string) *logrus.Logger {
	log := logrus.New()
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&logrus.JSONFormatter{})
	// log.SetFormatter(&logrus.TextFormatter{
	// 	DisableColors: true,
	// 	FullTimestamp: true,
	// })

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(outputFile)

	switch debugLevel {
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	default:
		log.SetLevel(logrus.DebugLevel)
	}
	return log
}
