package logging

import (
	"log"
	"os"
)

type Logging struct {
	File     *LogBackend
	Cmd      *LogBackend
	LogLevel Level
}

func NewLogging(filePath string, isShowCmd bool) (*Logging, error) {
	var err error
	var file, cmd *LogBackend

	if "" != filePath {
		f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm|os.ModeTemporary)

		if err != nil {
			panic(err)
		}

		file = NewLogBackend(f, "\r\n", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	}

	if isShowCmd {
		cmd = NewLogBackend(os.Stdout, "\r\n", log.Ldate|log.Lmicroseconds|log.Lshortfile)
		cmd.Color = true
	}

	return &Logging{
		File:     file,
		Cmd:      cmd,
		LogLevel: DEBUG,
	}, err
}

func (logging *Logging) SetOutPutLevel(strLevel string) {

	switch strLevel {
	case "CRITICAL":
		logging.LogLevel = CRITICAL
	case "ERROR":
		logging.LogLevel = ERROR
	case "WARNING":
		logging.LogLevel = WARNING
	case "NOTICE":
		logging.LogLevel = NOTICE
	case "INFO":
		logging.LogLevel = INFO
	case "DEBUG":
		logging.LogLevel = DEBUG
	}
}

func (logging *Logging) Critical(strMsg string) {
	logging.log(CRITICAL, 0, strMsg)
}

func (logging *Logging) Error(strMsg string) {
	logging.log(ERROR, 0, strMsg)
}

func (logging *Logging) Warning(strMsg string) {
	logging.log(WARNING, 0, strMsg)
}

func (logging *Logging) Info(strMsg string) {
	logging.log(INFO, 0, strMsg)
}

func (logging *Logging) Debug(strMsg string) {
	logging.log(DEBUG, 0, strMsg)
}

func (logging *Logging) log(level Level, calldepth int, strMsg string) {

	if level > logging.LogLevel {
		return
	}

	if logging.File != nil {
		logging.File.Log(level, calldepth, strMsg)
	}

	if logging.Cmd != nil {
		logging.Cmd.Log(level, calldepth, strMsg)
	}

}
