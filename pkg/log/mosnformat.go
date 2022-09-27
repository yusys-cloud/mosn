package log

import (
	"mosn.io/pkg/log"
)

var MosnLogger log.ErrorLogger

func init() {
	logger, err := GetOrCreateLogger("", nil)
	if err != nil {
		panic("init default logger error: " + err.Error())
	}
	MosnLogger = &SimpleErrorLog{
		Logger: logger,
		Level:  INFO,
	}
}

type SimpleErrorLog struct {
	*log.Logger
	Formatter func(lv string, alert string, format string) string
	Level     Level
}

func MosnFormatter(lv string, alert string, format string) string {
	lvstr01 := lv[1:len(lv)]
	lvstr02 := lvstr01[0 : len(lvstr01)-1]
	lv = lvstr02
	if alert == "" {
		return CacheTime() + " " + lv + " " + format
	}
	// return utils.CacheTime() + " " + lv + " [" + alert + "] " + format
	return CacheTime() + " " + lv + "+ alert +" + format
}

func (l *SimpleErrorLog) Alertf(alert string, format string, args ...interface{}) {
	if l.Disable() {
		return
	}
	if l.Level >= ERROR {
		var fs string
		if l.Formatter != nil {
			fs = l.Formatter(log.ErrorPre, alert, format)
		} else {
			fs = MosnFormatter(log.ErrorPre, alert, format)
		}
		l.Printf(fs, args...)
	}
}
func (l *SimpleErrorLog) levelf(lv string, format string, args ...interface{}) {
	if l.Disable() {
		return
	}
	fs := ""
	if l.Formatter != nil {
		fs = l.Formatter(lv, "", format)
	} else {
		fs = MosnFormatter(lv, "", format)
	}
	l.Printf(fs, args...)
}

func (l *SimpleErrorLog) Infof(format string, args ...interface{}) {
	if l.Level >= INFO {
		l.levelf(log.InfoPre, format, args...)
	}
}

func (l *SimpleErrorLog) Debugf(format string, args ...interface{}) {
	if l.Level >= DEBUG {
		l.levelf(log.DebugPre, format, args...)
	}
}

func (l *SimpleErrorLog) Warnf(format string, args ...interface{}) {
	if l.Level >= WARN {
		l.levelf(log.WarnPre, format, args...)
	}
}

func (l *SimpleErrorLog) Errorf(format string, args ...interface{}) {
	if l.Level >= ERROR {
		l.levelf(log.ErrorPre, format, args...)
	}
}

func (l *SimpleErrorLog) Tracef(format string, args ...interface{}) {
	if l.Level >= TRACE {
		l.levelf(log.TracePre, format, args...)
	}
}

func (l *SimpleErrorLog) Fatalf(format string, args ...interface{}) {
	var s string
	if l.Formatter != nil {
		s = l.Formatter(log.FatalPre, "", format)
	} else {
		s = MosnFormatter(log.FatalPre, "", format)
	}
	l.Logger.Fatalf(s, args...)
}

func (l *SimpleErrorLog) SetLogLevel(level Level) {
	l.Level = level
}
func (l *SimpleErrorLog) GetLogLevel() Level {
	return l.Level
}
