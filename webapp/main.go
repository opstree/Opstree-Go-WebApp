package webapp

import (
    "net/http"
    dbcheck "github.com/dimiro1/health/db"
    "github.com/dimiro1/health"
)

func Run() {
    generateLogging()
    createDatabaseTable()
    db := dbConn()
    mysql := dbcheck.NewMySQLChecker(db)
    handler := health.NewHandler()
    handler.AddChecker("MySQL", mysql)
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
