package main

import (
	"fmt"
	"log"
	"net/http"

	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/count", service.CounterHandler)
	http.HandleFunc("/api/GetObject", service.GetObjectHander)
	http.HandleFunc("/api/ObjectList", service.ObjectList)
	http.HandleFunc("/api/PutObject", service.PutObject)
	http.HandleFunc("/api/CreateObject", service.CreateObject)
	// http.HandleFunc("/api/UpdateObject", service.UpdateObject)
	// http.HandleFunc("/api/DeleteObject", service.DeleteObject)

	log.Fatal(http.ListenAndServe(":80", nil))

}
