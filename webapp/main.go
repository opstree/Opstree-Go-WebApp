package webapp

import (
	"github.com/dimiro1/health"
	dbcheck "github.com/dimiro1/health/db"
	"github.com/dimiro1/health/redis"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"net/http"
	"os"
	"time"
)

func Run() {

	propertyfile := "/etc/conf.d/ot-go-webapp/application.ini"
	if fileExists(propertyfile) {
		vaules, err := ini.Load(propertyfile)
		if err != nil {
			log.Error("No property file found in " + propertyfile)
		}
		redisHost = vaules.Section("redis").Key("REDIS_HOST").String()
		redisPort = vaules.Section("redis").Key("REDIS_PORT").String()
		logStdout()
		log.WithFields(log.Fields{
			"file": propertyfile,
		}).Info("Reading properties from " + propertyfile)
		logFile("access")
		log.WithFields(log.Fields{
			"file": propertyfile,
		}).Info("Reading properties from " + propertyfile)
	} else {
		redisHost = os.Getenv("REDIS_HOST")
		redisPort = os.Getenv("REDIS_PORT")
		logStdout()
		log.WithFields(log.Fields{
			"file": propertyfile,
		}).Info("No property file found, using environment variables")
		logFile("access")
		log.WithFields(log.Fields{
			"file": propertyfile,
		}).Info("No property file found, using environment variables")
	}

	generateLogsFile()
	time.Sleep(60 * time.Second)
	createDatabaseTable()
	db := dbConn()
	mysql := dbcheck.NewMySQLChecker(db)
	handler := health.NewHandler()
	handler.AddChecker("MySQL", mysql)
	handler.AddChecker("Redis", redis.NewChecker("tcp", redisHost+":"+redisPort))
	http.Handle("/health", handler)
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
