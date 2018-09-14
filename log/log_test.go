package tests

import (
	"flag"
	"github.com/1pb-club/golory/log"
)

var (
	confFilePath string
	defaultLog   *zap.Logger
)

func TestLog(t *testing.T) {
	flag.StringVar(&confFilePath, "c", "conf.toml", "配置文件路径")
	flag.Parse()

	golory.Init(confFilePath)
	defaultLog = golory.Log("defaultLog")
	defaultLog.Info("something")
}
