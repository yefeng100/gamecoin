package main

import (
	"fmt"
	"framework/log"
	"goclient/logic/msg"
	"goclient/modules/conf"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//初始化server配置
	_, err := conf.InitServer()
	if err != nil {
		fmt.Printf("init server failed, err:%v\n", err)
		return
	}
	//初始化日志
	logName := conf.GetIns().GetString("log.logName")
	logLevel := conf.GetIns().GetString("log.logLevel")
	log.InitLog(logName, logLevel)
	//处理事务
	msg.NewMsgMgr()
	//退出主携程
	notifyExit()
}

func destroyModules() {

}

func notifyExit() {
	sigOs := make(chan os.Signal, 1)
	signal.Notify(sigOs, os.Kill, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-sigOs
	fmt.Println("notifyExit--destroy--", sig.String())
	now := time.Now()
	//销毁模块
	destroyModules()
	fmt.Println("notifyExit--exit--", time.Since(now).String())
}
