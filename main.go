package main

import (
	"flag"
	"fmt"
	lo "log"
	"os"
	"path"
)

var (
	name = flag.String("name", "NoName", "Name of the peer")
	port = ""
	log  = lo.Default()
)

func main() {
	flag.Parse()

	prefix := fmt.Sprintf("%-8s: ", *name)
	logFileName := path.Join("logs", fmt.Sprintf("%s.log", *name))
	_ = os.Mkdir("logs", 0664)
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)

	if err != nil {
		log.Fatalf("Unable to open or create file %s: %v", logFileName, err)
	}

	log = lo.New(logFile, prefix, lo.Ltime)
}
