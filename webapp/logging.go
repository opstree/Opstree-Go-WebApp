package webapp

import (
	"os"
	"fmt"
	"io"
	log "github.com/sirupsen/logrus"
)

var accesslogfile = "/var/log/ot-go-webapp.access.log"
var errorlogfile = "/var/log/ot-go-webapp.error.log"
var databasefile = "/var/log/ot-go-webapp.database.log"
var infofile = "/var/log/ot-go-webapp.info.log"

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
    databasefile, err := os.OpenFile(databasefile, os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        fmt.Println(err)
    }
    defer errorfile.Close()
    infofile, err := os.OpenFile(infofile, os.O_CREATE|os.O_APPEND, 0644)
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
    } else if logtype == "error" {
        errorfile, err := os.OpenFile(errorlogfile, os.O_APPEND|os.O_WRONLY, 0644)
        if err != nil {
            fmt.Println(err)
        }
        log.SetOutput(errorfile)
    } else if logtype == "database" {
        databasefile, err := os.OpenFile(errorlogfile, os.O_APPEND|os.O_WRONLY, 0644)
        if err != nil {
            fmt.Println(err)
        }
        log.SetOutput(databasefile)
    } else {
        infofile, err := os.OpenFile(infofile, os.O_APPEND|os.O_WRONLY, 0644)
        if err != nil {
            fmt.Println(err)
        }
        log.SetOutput(infofile)
    }
}
