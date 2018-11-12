package logs

import (
	"fmt"

	"github.com/bigBarrage/roomManager/config"
)

const (
	_ = iota
	DISPLAY_ERROR_LEVEL_NOTICE
	DISPLAY_ERROR_LEVEL_WARNING
	DISPLAY_ERROR_LEVEL_ERROR
	DISPLAY_ERROR_LEVEL_FATAL

	GREEN  = "\033[;;32m"
	YELLOW = "\033[;;33m"
	PINK   = "\033[;;35m"
	RED    = "\033[;;31m"
	NONE   = "\033[0m"
)

func DisplayLog(level int64, message string) {
	switch level {
	case DISPLAY_ERROR_LEVEL_NOTICE:
		fmt.Fprintln(config.ErrorLogPath, GREEN+"[NOTICE]"+NONE+message)
	case DISPLAY_ERROR_LEVEL_WARNING:
		fmt.Fprintln(config.ErrorLogPath, YELLOW+"[WARNING]"+NONE+message)
	case DISPLAY_ERROR_LEVEL_ERROR:
		fmt.Fprintln(config.ErrorLogPath, PINK+"[ERROR]"+NONE+message)
	case DISPLAY_ERROR_LEVEL_FATAL:
		fmt.Fprintln(config.ErrorLogPath, RED+"[FATAL]"+NONE+message)
	}
}
