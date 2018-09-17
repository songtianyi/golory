package main

import (
	"flag"
	"github.com/1pb-club/golory/log"
	"github.com/1pb-club/golory"
)

var (
	confFilePath string
	defaultLog   *log.Logger
)

func main() {
	flag.StringVar(&confFilePath, "c", "conf.toml", "配置文件路径")
	flag.Parse()

	golory.Boot(confFilePath)
	defaultLog = golory.Log("defaultLog")
	defaultLog.Info("something")
}