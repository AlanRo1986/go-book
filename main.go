package main

import (
	"book/app"
	"log"
	"net/http"
	"time"
)

func init()  {

}

func main()  {
	//文件系统
	//fs := http.FileSystem(http.Dir("e:/tools"))
	//http.Handle("/", http.FileServer(fs))
	//log.Fatal(http.ListenAndServe(":8080", nil))

	port := "8080"
	web := http.Server{
		Addr:           ":"+port,
		Handler:        app.HttpHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	app.Router()

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)

	log.Fatal(web.ListenAndServe())
}
