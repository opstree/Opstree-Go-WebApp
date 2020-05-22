package webapp

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"net/http"
	"os"
	"text/template"
	"time"
)

type Employee struct {
	Id    int
	Name  string
	City  string
	Email string
	Date  string
}

var tmpl = template.Must(template.New("Employee Management Template").Parse(htmltemplate))

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbName := "employeedb"
	var dbUser string
	var dbPass string
	var dbUrl string
	var dbPort string
	propertyfile := "/etc/conf.d/ot-go-webapp/application.ini"

	if fileExists(propertyfile) {
		vaules, err := ini.Load(propertyfile)
		if err != nil {
			log.Error("No property file found in " + propertyfile)
		}
		dbUser = vaules.Section("database").Key("DB_USER").String()
		dbPass = vaules.Section("database").Key("DB_PASSWORD").String()
		dbUrl = vaules.Section("database").Key("DB_URL").String()
		dbPort = vaules.Section("database").Key("DB_PORT").String()
		logFile("access")
		log.WithFields(log.Fields{
			"file": propertyfile,
		}).Info("Reading properties from " + propertyfile)
	} else {
		dbUser = os.Getenv("DB_USER")
		dbPass = os.Getenv("DB_PASSWORD")
		dbUrl = os.Getenv("DB_URL")
		dbPort = os.Getenv("DB_PORT")
		logFile("access")
		log.WithFields(log.Fields{
			"file": propertyfile,
		}).Info("No property file found, using environment variables")
	}

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbUrl+":"+dbPort+")/"+dbName)

	if err != nil {
		logStdout()
		log.WithFields(log.Fields{
			"db_url": dbUrl,
		}).Error(err.Error())
		logFile("error")
		log.WithFields(log.Fields{
			"db_url": dbUrl,
		}).Error(err.Error())
	}
	return db
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func createDatabase() {
	db := dbConn()
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS employeedb")
	if err != nil {
		logStdout()
		log.WithFields(log.Fields{
			"database": "employeedb",
		}).Error(err.Error())
		logFile("error")
		log.WithFields(log.Fields{
			"database": "employeedb",
		}).Error(err.Error())
	} else {
		logStdout()
		log.WithFields(log.Fields{
			"database": "employeedb",
		}).Info("employeedb database is created")
		logFile("access")
		log.WithFields(log.Fields{
			"database": "employeedb",
		}).Info("employeedb database is created")
	}
	defer db.Close()
}

func createTable() {
	db := dbConn()

	_, err := db.Exec("USE employeedb")
	if err != nil {
		logStdout()
		log.WithFields(log.Fields{
			"database": "employeedb",
		}).Error(err.Error())
		logFile("error")
		log.WithFields(log.Fields{
			"database": "employeedb",
		}).Error(err.Error())
	} else {
		logStdout()
		log.WithFields(log.Fields{
			"database": "employeedb",
		}).Info("Using employeedb database")
		logFile("access")
		log.WithFields(log.Fields{
			"database": "employeedb",
		}).Info("Using employeedb database")
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Employee ( id int(6) NOT NULL AUTO_INCREMENT, name varchar(50) NOT NULL, city varchar(50) NOT NULL, email varchar(50) NOT NULL, date varchar(50), PRIMARY KEY (id) )")
	if err != nil {
		logStdout()
		log.WithFields(log.Fields{
			"table": "Employee",
		}).Error(err.Error())
		logFile("error")
		log.WithFields(log.Fields{
			"table": "Employee",
		}).Error(err.Error())
	} else {
		logStdout()
		log.WithFields(log.Fields{
			"table": "Employee",
		}).Info("Employee table is created in employeedb database")
		logFile("access")
		log.WithFields(log.Fields{
			"table": "Employee",
		}).Info("Employee table is created in employeedb database")
	}
	defer db.Close()
}

func createDatabaseTable() {
	createDatabase()
	createTable()
}

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	start := time.Now()
	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
	if err != nil {
		logStdout()
		log.WithFields(log.Fields{
			"query": "SELECT * FROM Employee ORDER BY id DESC",
		}).Error(err.Error())
		logFile("error")
		log.WithFields(log.Fields{
			"query": "SELECT * FROM Employee ORDER BY id DESC",
		}).Error(err.Error())
	}
	emp := Employee{}
	res := []Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		var email string
		var date string
		err = selDB.Scan(&id, &name, &city, &email, &date)
		if err != nil {
			logStdout()
			log.WithFields(log.Fields{
				"query": "scan &id, &name, &city, &email, &date",
			}).Error(err.Error())
			logFile("error")
			log.WithFields(log.Fields{
				"query": "scan &id, &name, &city, &email, &date",
			}).Error(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.Email = email
		emp.Date = date
		emp.City = city
		res = append(res, emp)
		logStdout()
		log.WithFields(log.Fields{
			"request_type":  "GET",
			"response_code": 200,
			"resonse_time":  time.Since(start),
			"request_url":   r.URL.Path,
		}).Info("Get request on index page")
		logFile("access")
		log.WithFields(log.Fields{
			"request_type":  "GET",
			"response_code": 200,
			"resonse_time":  time.Since(start),
			"request_url":   r.URL.Path,
		}).Info("Get request on index page")
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	start := time.Now()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
	if err != nil {
		logStdout()
		log.WithFields(log.Fields{
			"query": "SELECT * FROM Employee WHERE id",
			"id":    nId,
		}).Error(err.Error())
		logFile("error")
		log.WithFields(log.Fields{
			"query": "SELECT * FROM Employee WHERE id",
			"id":    nId,
		}).Error(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		var email string
		var date string
		err = selDB.Scan(&id, &name, &city, &email, &date)
		if err != nil {
			logStdout()
			log.WithFields(log.Fields{
				"query": "scan &id, &name, &city, &email, &date",
			}).Error(err.Error())
			logFile("error")
			log.WithFields(log.Fields{
				"query": "scan &id, &name, &city, &email, &date",
			}).Error(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.Email = email
		emp.Date = date
		emp.City = city
		logStdout()
		log.WithFields(log.Fields{
			"request_type":   "GET",
			"employee_name":  name,
			"employee_email": email,
			"employee_date":  date,
			"employee_city":  city,
			"response_code":  200,
			"resonse_time":   time.Since(start),
			"request_url":    r.URL.Path,
		}).Info("Get request on show page for " + name)
		logFile("access")
		log.WithFields(log.Fields{
			"request_type":   "GET",
			"employee_name":  name,
			"employee_email": email,
			"employee_date":  date,
			"employee_city":  city,
			"response_code":  200,
			"resonse_time":   time.Since(start),
			"request_url":    r.URL.Path,
		}).Info("Get request on show page for " + name)
	}
	tmpl.ExecuteTemplate(w, "Show", emp)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	start := time.Now()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
	if err != nil {
		logStdout()
		log.WithFields(log.Fields{
			"query": "SELECT * FROM Employee WHERE id",
			"id":    nId,
		}).Error(err.Error())
		logFile("error")
		log.WithFields(log.Fields{
			"query": "SELECT * FROM Employee WHERE id",
			"id":    nId,
		}).Error(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		var email string
		var date string
		err = selDB.Scan(&id, &name, &city, &email, &date)
		if err != nil {
			logStdout()
			log.WithFields(log.Fields{
				"query": "scan &id, &name, &city, &email, &date",
			}).Error(err.Error())
			logFile("error")
			log.WithFields(log.Fields{
				"query": "scan &id, &name, &city, &email, &date",
			}).Error(err.Error())
		}
		emp.Id = id
		emp.Date = date
		emp.Email = email
		emp.Name = name
		emp.City = city
		logStdout()
		log.WithFields(log.Fields{
			"request_type":   "POST",
			"employee_name":  name,
			"employee_email": email,
			"employee_date":  date,
			"employee_city":  city,
			"response_code":  200,
			"resonse_time":   time.Since(start),
			"request_url":    r.URL.Path,
		}).Info("Post request on edit page for " + name)
		logFile("access")
		log.WithFields(log.Fields{
			"request_type":   "POST",
			"employee_name":  name,
			"employee_email": email,
			"employee_date":  date,
			"employee_city":  city,
			"response_code":  200,
			"resonse_time":   time.Since(start),
			"request_url":    r.URL.Path,
		}).Info("Post request on edit page for " + name)
	}
	tmpl.ExecuteTemplate(w, "Edit", emp)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	start := time.Now()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		email := r.FormValue("email")
		date := r.FormValue("date")
		insForm, err := db.Prepare("INSERT INTO Employee(name, city, email, date) VALUES(?,?,?,?)")
		if err != nil {
			logStdout()
			log.WithFields(log.Fields{
				"query": "INSERT INTO Employee (name, city, email, date)",
			}).Error(err.Error())
			logFile("error")
			log.WithFields(log.Fields{
				"query": "INSERT INTO Employee (name, city, email, date)",
			}).Error(err.Error())
		}
		insForm.Exec(name, city, email, date)
		logStdout()
		log.WithFields(log.Fields{
			"request_type":   "POST",
			"employee_name":  name,
			"employee_email": email,
			"employee_date":  date,
			"employee_city":  city,
			"response_code":  200,
			"resonse_time":   time.Since(start),
			"request_url":    r.URL.Path,
		}).Info("Post request on insert page for " + name)
		logFile("access")
		log.WithFields(log.Fields{
			"request_type":   "POST",
			"employee_name":  name,
			"employee_email": email,
			"employee_date":  date,
			"employee_city":  city,
			"response_code":  200,
			"resonse_time":   time.Since(start),
			"request_url":    r.URL.Path,
		}).Info("Post request on insert page for " + name)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	start := time.Now()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")
		email := r.FormValue("email")
		date := r.FormValue("date")
		insForm, err := db.Prepare("UPDATE Employee SET name=?, city=?, email=?, date=? WHERE id=?")
		if err != nil {
			logStdout()
			log.WithFields(log.Fields{
				"query": "UPDATE Employee SET name=?, city=?, email=?, date=? WHERE id=?",
			}).Error(err.Error())
			logFile("error")
			log.WithFields(log.Fields{
				"query": "UPDATE Employee SET name=?, city=?, email=?, date=? WHERE id=?",
			}).Error(err.Error())
		}
		insForm.Exec(name, city, email, date, id)
		logStdout()
		log.WithFields(log.Fields{
			"request_type":   "POST",
			"employee_name":  name,
			"employee_email": email,
			"employee_date":  date,
			"employee_city":  city,
			"response_code":  200,
			"resonse_time":   time.Since(start),
			"request_url":    r.URL.Path,
		}).Info("Post request on update page for " + name)
		logFile("access")
		log.WithFields(log.Fields{
			"request_type":   "POST",
			"employee_name":  name,
			"employee_email": email,
			"employee_date":  date,
			"employee_city":  city,
			"response_code":  200,
			"resonse_time":   time.Since(start),
			"request_url":    r.URL.Path,
		}).Info("Post request on update page for " + name)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	start := time.Now()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
	if err != nil {
		logStdout()
		log.WithFields(log.Fields{
			"query": "DELETE FROM Employee WHERE id",
			"id":    emp,
		}).Error(err.Error())
		logFile("error")
		log.WithFields(log.Fields{
			"query": "DELETE FROM Employee WHERE id",
			"id":    emp,
		}).Error(err.Error())
	}
	delForm.Exec(emp)
	logStdout()
	log.WithFields(log.Fields{
		"request_type":  "POST",
		"id":            emp,
		"response_code": 200,
		"resonse_time":  time.Since(start),
		"request_url":   r.URL.Path,
	}).Info("Post request on delete page for " + emp)
	logFile("access")
	log.WithFields(log.Fields{
		"request_type":  "POST",
		"id":            emp,
		"response_code": 200,
		"resonse_time":  time.Since(start),
		"request_url":   r.URL.Path,
	}).Info("Post request on delete page for " + emp)
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}
