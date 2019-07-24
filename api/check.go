package api

import (
	"execworker/runner"
	"github.com/valyala/fasthttp"
)

//本接口用来监控有几个任务正在运行

func Check(ctx *fasthttp.RequestCtx) string {

	str := runner.Info()
	return str
}