package main

import (

	"fmt"
	"io"
	"time"
	"github.com/rz1226/simplegokit/httpkit"
	"strings"
)


//发送任务到接口测试

var kit *httpkit.HttpClient
func init(){
	kit  = httpkit.NewHttpClient(10,10)
}

func main() {
	for i:=0;i<100 ;i++  {
		go testx()
	}
	time.Sleep(time.Second*100)
}

func testx(){
	fmt.Println("starting")
	str := `["php","/mnt/win/vmfiles/goproject/execworker/test.php"]`
	//str := `["ls"]`

	url := "http://127.0.0.1/run"
	var buf io.Reader
	buf = strings.NewReader("job="+ str )
	str, err := kit.Post(url, "application/x-www-form-urlencoded;charset=utf-8", buf )
	fmt.Println(str, err )


}


