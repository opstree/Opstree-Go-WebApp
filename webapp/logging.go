package webapp

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

var accesslogfile = "/var/log/ot-go-webapp.access.log"
var errorlogfile = "/var/log/ot-go-webapp.error.log"

func generateLogsFile() {
	accessfile, err := os.OpenFile(accesslogfile, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer accessfile.Close()
	errorfile, err := os.OpenFile(errorlogfile, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer errorfile.Close()
}

func logStdout() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
	})
	mw := io.MultiWriter(os.Stdout)
	log.SetOutput(mw)
}

func logFile(logtype string) {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
	})
	if logtype == "access" {
		accessfile, err := os.OpenFile(accesslogfile, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}
		log.SetOutput(accessfile)
	} else {
		errorfile, err := os.OpenFile(errorlogfile, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}
		log.SetOutput(errorfile)
	}
}
