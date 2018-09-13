package main

import (
	"flag"
	"github.com/1pb-club/golory"
	"go.uber.org/zap"
)

var (
	confFilePath string
	defaultLog   *zap.Logger
)

func main() {
	flag.StringVar(&confFilePath, "c", "conf.toml", "配置文件路径")
	flag.Parse()

	golory.Init(confFilePath)
	defaultLog = golory.Log("defaultLog")
	defaultLog.Info("something")
}
