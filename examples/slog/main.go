package main

import (
	"github.com/coinrust/log"
)

func init() {
	log.Init("./test",
		log.DebugLevel,
		log.SetMaxFileSize(10),
		log.SetMaxBackups(2),
		log.SetCaller(true),
		log.SetStdout(true),
		log.SetCompress(true),
		log.SetSLog(true))
}

func main() {
	defer log.Sync()
	for i := int64(0); i < 1000000; i++ {
		log.Infof("hello %v %v", "U", i)
	}
}
