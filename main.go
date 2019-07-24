package main

import(
	"execworker/api"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"fmt"
	"log"

)


func main() {


	fmt.Println("starting")
	router := fasthttprouter.New()
	router.GET("/check", api.GenApifunc(api.Check, "check"))
	router.POST("/run",api.GenApifunc(api.Run, "run"))
	log.Fatal(fasthttp.ListenAndServe(":80", router.Handler))
}
