/*
	Copyright (c) 2014 Ray Qiu <ray.qiu@gmail.com>
 	All rights reserved.

	Version: 	0.1
	Authors: 	Ray Qiu <ray.qiu@gmail.com>
	Date:    	April, 2014

	A log trace enabler that only display logging messages for relevant levels

*/

package logtrace

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

const (
	LOG_FATAL = iota
	LOG_ERROR
	LOG_WARNING
	LOG_INFO
	LOG_TRACE
)

var (
	once  sync.Once
	_self logtrace
)

type logtrace struct {
	mutex    sync.Mutex
	logLevel int
}

func init() {
	once.Do(func() {
		_self := logtrace{}
		_ = _self // so no compile warning
	})
	_self.logLevel = LOG_WARNING
}

func PrintSelf() {
	log.Println(_self)
}

func SetLogLevel(l int) error {
	if l < LOG_FATAL || l > LOG_TRACE {
		errString := fmt.Sprintf("logLevel must be between %d and %d!", LOG_TRACE, LOG_FATAL)
		return errors.New(errString)
	}
	_self.mutex.Lock()
	_self.logLevel = l
	_self.mutex.Unlock()
	return nil
}

func LogTraceln(logger *log.Logger, v ...interface{}) {
	if _self.logLevel < LOG_TRACE {
		return
	}
	logger.Println(v...)
}

func LogTracef(logger *log.Logger, format string, v ...interface{}) {
	if _self.logLevel < LOG_TRACE {
		return
	}
	logger.Printf(format, v...)
}

func LogInfoln(logger *log.Logger, v ...interface{}) {
	if _self.logLevel < LOG_INFO {
		return
	}
	logger.Println(v...)
}

func LogInfof(logger *log.Logger, format string, v ...interface{}) {
	if _self.logLevel < LOG_INFO {
		return
	}
	logger.Printf(format, v...)
}

func LogWarningln(logger *log.Logger, v ...interface{}) {
	if _self.logLevel < LOG_WARNING {
		return
	}
	logger.Println(v...)
}

func LogWarningf(logger *log.Logger, format string, v ...interface{}) {
	if _self.logLevel < LOG_WARNING {
		return
	}
	logger.Printf(format, v...)
}

func LogErrorln(logger *log.Logger, v ...interface{}) {
	if _self.logLevel < LOG_ERROR {
		return
	}
	logger.Println(v...)
}

func LogErrorf(logger *log.Logger, format string, v ...interface{}) {
	if _self.logLevel < LOG_ERROR {
		return
	}
	logger.Printf(format, v...)
}

func LogFatalln(logger *log.Logger, v ...interface{}) {
	if _self.logLevel < LOG_FATAL {
		return
	}
	logger.Println(v...)
}

func LogFatalf(logger *log.Logger, format string, v ...interface{}) {
	if _self.logLevel < LOG_FATAL {
		return
	}
	logger.Printf(format, v...)
}
