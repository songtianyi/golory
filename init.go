package golory

func Init(configFile string) {
	invokerObj := initInvoker(configFile)
	invokerObj.initLog()
}
