package webapp

import (
    "database/sql"
    log "github.com/sirupsen/logrus"
    "fmt"
    "io"
    "gopkg.in/ini.v1"
    "os"
    "net/http"
    "text/template"
    _ "github.com/go-sql-driver/mysql"
)

type Employee struct {
    Id    int
    Name  string
    City  string
    Email string
    Date string
}

var accesslogfile = "/var/log/ot-go-webapp.access.log"
var errorlogfile = "/var/log/ot-go-webapp.error.log"

var tmpl = template.Must(template.New("Employee Management Template").Parse(htmltemplate))

func generateLogging() {
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

func loggingInit() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
    })
    mw := io.MultiWriter(os.Stdout)
    log.SetOutput(mw)
}

func loggingLogFileInit(logtype string) {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
    })
    if logtype == "access" {
        accessfile, err := os.OpenFile(accesslogfile, os.O_APPEND|os.O_WRONLY, 0644)
        if err != nil {
            fmt.Println(err)
        }
        // defer accessfile.Close()
        log.SetOutput(accessfile)
    } else {
        errorfile, err := os.OpenFile(errorlogfile, os.O_APPEND|os.O_WRONLY, 0644)
        if err != nil {
            fmt.Println(err)
        }
        // defer errorfile.Close()
        log.SetOutput(errorfile)
    }
}

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbName := "employeedb"
    var dbUser string
    var dbPass string
    var dbUrl string
    var dbPort string
    propertyfile := "/etc/conf.d/ot-go-webapp/application.ini"

    if fileExists(propertyfile) {
        loggingInit()
        vaules, err := ini.Load(propertyfile)
        if err != nil {
            log.Error("No property file found in " + propertyfile)
        }
        dbUser = vaules.Section("database").Key("DB_USER").String()
        dbPass = vaules.Section("database").Key("DB_PASSWORD").String()
        dbUrl  = vaules.Section("database").Key("DB_URL").String()
        dbPort = vaules.Section("database").Key("DB_PORT").String()
        log.Info("Reading properties file " + propertyfile)
        loggingLogFileInit("access")
        log.Info("Reading properties file " + propertyfile)
    } else {
        dbUser = os.Getenv("DB_USER")
        dbPass = os.Getenv("DB_PASSWORD")
        dbUrl  = os.Getenv("DB_URL")
        dbPort = os.Getenv("DB_PORT")
        loggingInit()
        log.Info("No property file found, using environment variables")
        loggingLogFileInit("access")
        log.Info("No property file found, using environment variables")
    }

    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbUrl+":"+dbPort+")/"+dbName)

    if err != nil {
        loggingInit()
        log.Error(err.Error())
        loggingLogFileInit("error")
        log.Error(err.Error())
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
        loggingInit()
        log.Error(err.Error())
        loggingLogFileInit("error")
        log.Error(err.Error())
	} else {
        loggingInit()
        log.Info("Database is created with name employeedb")
        loggingLogFileInit("access")
        log.Info("Database is created with name employeedb")
    }
    defer db.Close()
}

func createTable() {
    db := dbConn()

	_,err := db.Exec("USE employeedb")
	if err != nil {
        loggingInit()
        log.Error(err.Error())
        loggingLogFileInit("error")
        log.Error(err.Error())
	} else {
        loggingInit()
        log.Info("Using employeedb database")
        loggingLogFileInit("access")
        log.Info("Using employeedb database")
    }
    
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Employee ( id int(6) NOT NULL AUTO_INCREMENT, name varchar(50) NOT NULL, city varchar(50) NOT NULL, email varchar(50) NOT NULL, date varchar(50), PRIMARY KEY (id) )")
	if err != nil {
        loggingInit()
        log.Error(err.Error())
        loggingLogFileInit("error")
        log.Error(err.Error())
	} else {
        loggingInit()
        log.Info("Table is created with name Employee")
        loggingLogFileInit("access")
        log.Info("Table is created with name Employee")
    }
    defer db.Close()
}

func createDatabaseTable() {
    createDatabase()
    createTable()
}

func Index(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
    if err != nil {
        loggingInit()
        log.Error(err.Error())
        loggingLogFileInit("error")
        log.Error(err.Error())
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
            loggingInit()
            log.Error(err.Error())
            loggingLogFileInit("error")
            log.Error(err.Error())
        }
        emp.Id = id
        emp.Name = name
        emp.Email = email
        emp.Date = date
        emp.City = city
        res = append(res, emp)
        loggingInit()
        log.Info("Get request on the /index page")
        loggingLogFileInit("access")
        log.Info("Get request on the /index page")
    }
    tmpl.ExecuteTemplate(w, "Index", res)
    defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
    if err != nil {
        loggingInit()
        log.Error(err.Error())
        loggingLogFileInit("error")
        log.Error(err.Error())
    }
    emp := Employee{}
    for selDB.Next() {
        var id int
        var name, city string
        var email string
        var date string
        err = selDB.Scan(&id, &name, &city, &email, &date)
        if err != nil {
            loggingInit()
            log.Error(err.Error())
            loggingLogFileInit("error")
            log.Error(err.Error())
        }
        emp.Id = id
        emp.Name = name
        emp.Email = email
        emp.Date = date
        emp.City = city
        loggingInit()
        log.Info("Get request on the /show for " + emp.Name)
        loggingLogFileInit("access")
        log.Info("Get request on the /show for " + emp.Name)
    }
    tmpl.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
    if err != nil {
        loggingInit()
        log.Error(err.Error())
        loggingLogFileInit("error")
        log.Error(err.Error())
    }
    emp := Employee{}
    for selDB.Next() {
        var id int
        var name, city string
        var email string
        var date string
        err = selDB.Scan(&id, &name, &city, &email, &date)
        if err != nil {
            loggingInit()
            log.Error(err.Error())
            loggingLogFileInit("error")
            log.Error(err.Error())
        }
        emp.Id = id
        emp.Date = date
        emp.Email = email
        emp.Name = name
        emp.City = city
        loggingInit()
        log.Info("Post request on the /edit for " + emp.Name)
        loggingLogFileInit("access")
        log.Info("Post request on the /edit for " + emp.Name)
    }
    tmpl.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        city := r.FormValue("city")
        email := r.FormValue("email")
        date := r.FormValue("date")
        insForm, err := db.Prepare("INSERT INTO Employee(name, city, email, date) VALUES(?,?,?,?)")
        if err != nil {
            loggingInit()
            log.Error(err.Error())
            loggingLogFileInit("error")
            log.Error(err.Error())
        }
        insForm.Exec(name, city, email, date)
        loggingInit()
        log.Info("Post request on the /insert for " + name)
        loggingLogFileInit("access")
        log.Info("Post request on the /insert for " + name)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        city := r.FormValue("city")
        id := r.FormValue("uid")
        email := r.FormValue("email")
        date := r.FormValue("date")
        insForm, err := db.Prepare("UPDATE Employee SET name=?, city=?, email=?, date=? WHERE id=?")
        if err != nil {
            loggingInit()
            log.Error(err.Error())
            loggingLogFileInit("error")
            log.Error(err.Error())
        }
        insForm.Exec(name, city, email, date, id)
        loggingInit()
        log.Info("Post request on the /update for " + name)
        loggingLogFileInit("access")
        log.Info("Post request on the /update for " + name)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")
    delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
    if err != nil {
        loggingInit()
        log.Error(err.Error())
        loggingLogFileInit("error")
        log.Error(err.Error())
    }
    delForm.Exec(emp)
    loggingInit()
    log.Info("Post request on the /delete for " + emp)
    loggingLogFileInit("access")
    log.Info("Post request on the /delete for " + emp)
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}
