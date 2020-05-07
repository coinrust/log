package main

import "github.com/coinrust/log"

func init() {
	log.Init("./test.log",
		log.DebugLevel,
		log.SetCaller(true),
		log.SetStdout(true))
}

func main() {
	defer log.Sync()

	log.Infof("%v", "hello")
}
