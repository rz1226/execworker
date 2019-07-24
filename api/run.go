package api

import (
	"encoding/json"
	"execworker/runner"
	"fmt"
	"github.com/valyala/fasthttp"
	"os/exec"
)

//接收外部传过来的任务，并运行
func Run(ctx *fasthttp.RequestCtx) string {
	job := ctx.FormValue("job")

	var result []interface{}
	err := json.Unmarshal(job, &result)
	if err != nil {
		return API_JOBERR
	}
	fmt.Println(result )
	if len(result) == 0 {
		return API_JOBERR
	}
	command ,ok := result[0].(string)
	if !ok{
		return API_JOBERR
	}
	var args []interface{}

	if len(result)>1{
		args = result[1:]
	}

	var argsStr []string
	for _, v := range args{
		argsStr = append( argsStr, fmt.Sprint( v ))
	}


	cmd := exec.Command( command, argsStr... )
	go run(cmd)
	return API_OK
}

func run(cmd *exec.Cmd){
	x := runner.NewRunner(cmd)
	x.Run()
}