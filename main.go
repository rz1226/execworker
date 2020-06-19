package main

import (
	"execworker/api"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {

	fmt.Println("starting")
	router := fasthttprouter.New()
	router.GET("/check", api.GenApifunc(api.Check, "check"))
	router.POST("/run", api.GenApifunc(api.Run, "run"))
	log.Fatal(fasthttp.ListenAndServe(":80", router.Handler))
}
