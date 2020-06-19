package api

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

//本包的公用函数

const (
	API_OK     = `{"code":0,"msg":"ok"}`
	API_PANIC  = `{"code":40,"msg":"server error"}`
	API_JOBERR = `{"code":41,"msg":"json error"}`
)

func GenApifunc(f func(*fasthttp.RequestCtx) string, name string) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if co := recover(); co != nil {
				fmt.Fprint(ctx, API_PANIC, co, "api=", name)
			}
		}()
		str := f(ctx)
		fmt.Fprint(ctx, str)
	}
}

//生成json
func GenResultJson(code int, msg string) string {
	return `{"code":` + fmt.Sprint(code) + `,"msg":"` + msg + `"}`
}
