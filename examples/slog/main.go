package main

import (
	"github.com/coinrust/log"
)

func init() {
	log.Init("./test",
		log.DebugLevel,
		log.SetCaller(true),
		log.SetStdout(true),
		log.SetSLog(true))
}

func main() {
	defer log.Sync()

	log.Infof("hello %v %v", "U", 1)
}
