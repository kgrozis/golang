package main

import (
	"fmt"
	"log"
	"log/syslog"
)

// log.fatal() function is used to exit program as quickly as possible
//   It represents a major issue has occured

func main() {
	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "Some Program")

	if err != nil {
		log.Fatal(err) // exit status 1
	} else {
		log.SetOutput(sysLog)
	}

	log.Fatal(sysLog)
	fmt.Println("Will you see this?")

}
