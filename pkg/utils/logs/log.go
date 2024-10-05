package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type SysLog struct {
	File   *os.File
	StrDir string
}

func (sysLog *SysLog) Init() (*SysLog, error) {
	logFileName := time.Now().Format("2006-01-02") + ".log"
	logPath := filepath.Join(sysLog.StrDir, logFileName)

	var err error
	sysLog.File, err = os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	sysLog.File.WriteString(fmt.Sprintf("[%s]:[Info] Logger Initial!\n", time.Now().Format("2006-01-02 15:04:05")))

	return &SysLog{
		File: sysLog.File,
	}, nil
}

func (sysLog *SysLog) Write(sType string, format string, args ...interface{}) error {

	if sysLog.File == nil {
		return fmt.Errorf("log file is not initialized")
	}
	_, fullPath, line, ok := runtime.Caller(1)
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}

	file := filepath.Base(fullPath)

	_, err := sysLog.File.WriteString(fmt.Sprintf("[%s]:[%s]:Line %d: [%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), file, line, sType, fmt.Sprintf(format, args...)))
	if err != nil {
		return fmt.Errorf("error writing to log file: %v", err)
	}

	return nil
}

func (sysLog *SysLog) Close() error {
	if sysLog.File != nil {
		err := sysLog.File.Close()
		sysLog.File = nil
		return err
	}
	return nil
}
