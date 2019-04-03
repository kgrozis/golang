package main

import (
	"fmt"
	"log"        // allows you to send log msgs to syslog on unix
	"log/syslog" // defines logging level and logging facility
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])
	// first parameter of syslog.New is priority (logging facility + logging level)
	syslog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)

	if err != nil { // checks error value from syslog.New call
		log.Fatal(err)
	} else {
		log.SetOutput(syslog) // No Errors.  Set output to default logger
	}
	log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go!") // send info to logger

	log.Println("LOG_MAIL: Logging in Go!") // to syslog
	fmt.Println("Will you see this?")       // to stdout

}
