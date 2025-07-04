package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

type MyLog struct{}

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

func (MyLog) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue // info
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	// 自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		// 自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		// 自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	}
	return b.Bytes(), nil
}
func InitLogger() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(MyLog{})
	//logrus.SetFormatter(&logrus.JSONFormatter{})  // 产生json格式，方便向外发送
	logrus.AddHook(&MyHook{
		logPath: "logs",
	})
}

type MyHook struct {
	file    *os.File // 当前打开的日志文件
	errFile *os.File // 当前的错误日志文件
	date    string   // 当前日志的时间
	logPath string   // 日志的目录
}

func (hook *MyHook) Fire(entry *logrus.Entry) error {
	// *MyHook使用指针是因为需要修改结构体中的内容，比如hook.date = date

	// 1、写入到文件
	// 2、按时间分片
	// 3、错误的日志单独存放
	msg, _ := entry.String()                // 拿日志内容(你好)，string类型
	date := entry.Time.Format("2006-01-02") // 当前日期
	if hook.date != date {
		// 换时间
		hook.rotateFile(date)
		hook.date = date
	}
	if entry.Level <= logrus.ErrorLevel {
		hook.errFile.Write([]byte(msg))
	}
	hook.file.Write([]byte(msg))
	return nil
}
func (hook *MyHook) rotateFile(time string) error {
	if hook.file != nil {
		hook.file.Close()
	}
	if hook.file == nil {
		// 创建目录
		logDir := fmt.Sprintf("%s/%s", hook.logPath, time)
		os.MkdirAll(logDir, 0666)
		logPath := fmt.Sprintf("%s/info.log", logDir)
		file, _ := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		hook.file = file

		errLogPath := fmt.Sprintf("%s/err.log", logDir)
		errFile, _ := os.OpenFile(errLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		hook.errFile = errFile
	}
	return nil
}
func (*MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
