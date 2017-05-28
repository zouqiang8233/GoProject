package logging_test

import (
	"base/logging"
	"testing"
)

func TestCmd(t *testing.T) {
	log, err := logging.NewLogging("", true)

	if err == nil {
		log.Critical("12345")
		log.Error("12345")
		log.Warning("12345")
		log.Info("12345")
		log.Debug("12345")
	}
}

func TestFile(t *testing.T) {
	log, err := logging.NewLogging("C:\\test.txt", false)

	if err == nil {
		log.SetOutPutLevel("CRITICAL")
		log.Critical("12345")
		log.Error("12345")
		log.Warning("12345")
		log.Info("12345")
		log.Debug("12345")
	}
}

func TestFileAndCmd(t *testing.T) {
	log, err := logging.NewLogging("C:\\test.txt", true)

	if err == nil {
		log.Critical("12345")
		log.Error("12345")
		log.Warning("12345")
		log.Info("12345")
		log.Debug("12345")
	}
}
