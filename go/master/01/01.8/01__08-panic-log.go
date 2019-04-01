package main

import (
	"fmt"
	"log"
	"log/syslog"
)

// log.panic() function provides as much info about the failure as possible

func main() {
	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "Some Program")

	if err != nil {
		log.Panic(sysLog) // provide failure info to stdout
	} else {
		log.SetOutput(sysLog)
	}

	log.Panic(sysLog)
	fmt.Println("Will you see this?")

}
