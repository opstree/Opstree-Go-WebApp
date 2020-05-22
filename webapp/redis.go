package webapp

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"os"
	"strconv"
	"time"
	// "json"
	"github.com/gomodule/redigo/redis"
	"net/http"
)

var pool *redis.Pool

var redisHost string
var redisPort string

func initializeCache() *redis.Pool {
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

	return &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisHost+":"+redisPort)
		},
	}
}

func redisIndex(w http.ResponseWriter, r *http.Request) {
	pool = initializeCache()
	conn := pool.Get()
	emp := Employee{}
	res := []Employee{}
	keys_list, err := redis.Strings(conn.Do("KEYS", "*"))
	if err != nil {
		log.Error(err)
	}

	for _, key := range keys_list {
		id := covertString(key)
		emp.Id = id
		emp.Name = getRedisKey(key, "name")
		emp.Email = getRedisKey(key, "email")
		emp.Date = getRedisKey(key, "date")
		emp.City = getRedisKey(key, "city")
		res = append(res, emp)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
}

func redisUserShow(w http.ResponseWriter, r *http.Request) {
	nId := r.FormValue("id")
	emp := Employee{}

	id := covertString(nId)
	emp.Id = id
	emp.Name = getRedisKey("nId", "name")
	emp.Email = getRedisKey("nId", "email")
	emp.Date = getRedisKey("nId", "date")
	emp.City = getRedisKey("nId", "city")

	tmpl.ExecuteTemplate(w, "Show", emp)
}

func redisEditUser(w http.ResponseWriter, r *http.Request) {
	nId := r.FormValue("id")
	emp := Employee{}

	id := covertString(nId)
	emp.Id = id
	emp.Name = getRedisKey("nId", "name")
	emp.Email = getRedisKey("nId", "email")
	emp.Date = getRedisKey("nId", "date")
	emp.City = getRedisKey("nId", "city")

	tmpl.ExecuteTemplate(w, "Edit", emp)
}

func redisInsertUser(w http.ResponseWriter, r *http.Request) {
	pool = initializeCache()
	conn := pool.Get()
	if r.Method == "POST" {
		nId := r.FormValue("id")
		name := r.FormValue("name")
		city := r.FormValue("city")
		email := r.FormValue("email")
		date := r.FormValue("date")

		insForm, err := conn.Do("HMSET", nId, "name", name, "email", email, "date", date, "city", city)
		if err != nil {
			log.Error(err)
		}
		fmt.Print(insForm)
	}
	http.Redirect(w, r, "/", 301)
}

func redisUpdateUser(w http.ResponseWriter, r *http.Request) {
	pool = initializeCache()
	conn := pool.Get()
	if r.Method == "POST" {
		nId := r.FormValue("id")
		name := r.FormValue("name")
		city := r.FormValue("city")
		email := r.FormValue("email")
		date := r.FormValue("date")
		insForm, err := conn.Do("HMSET", nId, "name", name, "email", email, "date", date, "city", city)
		if err != nil {
			log.Error(err)
		}
		fmt.Print(insForm)
	}
	http.Redirect(w, r, "/", 301)
}

func redisDeleteUser(w http.ResponseWriter, r *http.Request) {
	pool = initializeCache()
	conn := pool.Get()
	nId := r.FormValue("id")
	insForm, err := conn.Do("DEL", nId)
	if err != nil {
		log.Error(err)
	}
	fmt.Print(insForm)
	fmt.Println(insForm)
	http.Redirect(w, r, "/", 301)
}

func getRedisKey(key string, value string) string {
	pool = initializeCache()
	conn := pool.Get()
	key, err := redis.String(conn.Do("HGET", key, value))
	if err != nil {
		log.Error(err)
	}
	return key
}

func covertString(key string) int {
	data, err := strconv.Atoi(key)
	if err != nil {
		log.Error(err)
	}
	return data
}
