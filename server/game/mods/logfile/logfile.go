package logfile

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/topfreegames/pitaya/v2/logger"
	"github.com/topfreegames/pitaya/v2/logger/interfaces"
	logruswrapper "github.com/topfreegames/pitaya/v2/logger/logrus"
	"io"
	"os"
	"time"
)

var gSvType = ""      //保存全局类型,服务器
var gLogFile *os.File //上一个文件 句柄

// LogLoad 加载日志
func LogLoad(svType string, logLevel int, logFileDate bool) {
	gSvType = svType
	folderName := "logs"
	//创建文件夹
	_ = os.Mkdir("./"+folderName, os.ModePerm)
	_ = os.Mkdir("./"+folderName+"/"+svType, os.ModePerm)

	logger.SetLogger(getNewLogger(folderName, logLevel, logFileDate))
	logTicker := time.NewTicker(10 * time.Hour)
	go func() {
		for range logTicker.C {
			logger.SetLogger(getNewLogger(folderName, logLevel, logFileDate))
		}
	}()
}

func getNewLogger(folderName string, logLevel int, logFileDate bool) interfaces.Logger {
	plog := logrus.New()
	plog.Formatter = &logrus.JSONFormatter{}
	plog.SetReportCaller(true) //输出方法名
	plog.Formatter = &logrus.TextFormatter{
		DisableColors:   false,         // 关闭控制台颜色输出
		TimestampFormat: time.DateTime, //输出日志时间格式
	}
	plog.Level = logrus.Level(logLevel)

	filename := ""
	if logFileDate {
		filename = "./" + folderName + "/" + gSvType + "/" + gSvType + "_" + time.Now().Format("2006-01-02") + ".txt"
	} else {
		filename = "./" + folderName + "/" + gSvType + "/" + gSvType + ".txt"
	}
	fmt.Println("logifle=%s" + filename)
	var f *os.File

	if checkFileIsExist(filename) { //如果文件存在
		f, _ = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		//fmt.Println("log文件存在")
	} else {
		f, _ = os.Create(filename) //创建文件
		//fmt.Println("log文件不存在")
	}
	plog.SetOutput(io.MultiWriter(os.Stdout, f))
	log := plog.WithFields(logrus.Fields{
		"source": "pitaya",
	})
	if gLogFile != nil {
		_ = gLogFile.Close()
		gLogFile = nil
	}
	gLogFile = f
	return logruswrapper.NewWithFieldLogger(log)
}

// 文件存在
func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
