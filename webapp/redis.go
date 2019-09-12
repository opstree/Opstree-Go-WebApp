package webapp

import (
    log "github.com/sirupsen/logrus"
    "fmt"
	// "io"
	"time"
    "gopkg.in/ini.v1"
	"os"
	"github.com/gomodule/redigo/redis"	
)

var pool *redis.Pool

func initializeCache() *redis.Pool {
    var redisHost string
	var redisPort string
	propertyfile := "/etc/conf.d/ot-go-webapp/application.ini"

    if fileExists(propertyfile) {
        loggingInit()
        vaules, err := ini.Load(propertyfile)
        if err != nil {
            log.Error("No property file found in " + propertyfile)
        }
        redisHost = vaules.Section("redis").Key("REDIS_HOST").String()
        redisPort = vaules.Section("redis").Key("REDIS_PORT").String()
        log.Info("Reading properties file " + propertyfile)
    } else {
        redisHost = os.Getenv("REDIS_HOST")
        redisPort = os.Getenv("REDIS_PORT")
        loggingInit()
        log.Info("No property file found, using environment variables")
	}

	return &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisHost + ":" + redisPort)
		},
	}
}

func redisIndex() {
	pool = initializeCache()
	conn :=  pool.Get()
	keys_list, err := redis.Strings(conn.Do("KEYS", "1"))
	if err != nil {
		log.Error(err)
	}
	for _, key := range keys_list {
		reply, err := redis.StringMap(conn.Do("HGETALL", key))
		if err != nil {
			log.Error(err)
		}
		for key, value := range reply {
			fmt.Println("key is %s, value is %s", key, value)
		}
	}
}
