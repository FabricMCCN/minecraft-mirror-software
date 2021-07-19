package logger

import (
	"fmt"
	"time"
)

const (
	infoString = "%s [INF]:%s\n"
	logString  = "%s [LOG]:%s\n"
	warnString = "%s [WAR]:%s\n"
	erroString = "%s [ERR]:%s\n"
)

func Info(string2 string) {
	fmt.Printf(infoString, timeNow(), string2)
}

func Infof(string2 string, interface2 ...interface{}) {
	fmt.Printf(infoString, timeNow(), fmt.Sprintf(string2, interface2...))
}

func Error(err error) error {
	if err != nil {
		fmt.Printf(erroString, timeNow(), err.Error())
	}
	return err
}

func Warn(string2 string) {
	fmt.Printf(warnString, timeNow(), string2)
}

func Warnf(string2 string, interface2 ...interface{}) {
	fmt.Printf(warnString, timeNow(), fmt.Sprintf(string2, interface2...))
}

func timeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
