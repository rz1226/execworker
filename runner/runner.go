package runner

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/rz1226/utilx2/randx"
	"time"
)

//用来执行脚本任务  执行的时候放入runnerlist列表，完毕从里面删除。放入的目的是监控同时有多少任务进行

type Runner struct {
	id               string
	startTime        string
	cmd              *exec.Cmd
	notifyChanFinish chan struct{}
	outBuf           *CmdOutBuffer
}

func NewRunner(cmd *exec.Cmd) *Runner {
	r := &Runner{}
	r.cmd = cmd
	r.notifyChanFinish = make(chan struct{})
	r.outBuf = newCmdOutBuffer()
	r.cmd.Stdout = r.outBuf
	r.cmd.Stderr = r.outBuf
	r.id = strings.ToLower(randx.GetRandomString(10))
	r.startTime = time.Now().Format("2006-01-02 15:04:05")
	return r
}

func (r *Runner) Run() {
	RunnerAll.Add(r)
	go r.readCmdOutStrings()
	err := r.cmd.Run()
	if err != nil {
		//log.Fatal(err)
		r.outBuf.Write([]byte("出问题了，以下是错误信息:" + err.Error()))
		r.outBuf.Write([]byte("\n"))
	}
	//跑完通知

	r.notifyChanFinish <- struct{}{}
}

func (r *Runner) readCmdOutStrings() {
	defer func() {
		fmt.Println(r.id, "exit.....")
	}()

	for {
		time.Sleep(time.Microsecond * 10)
		str, err := r.outBuf.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				time.Sleep(time.Microsecond * 1000)
			} else {

			}

		} else {
			fmt.Print("<"+r.id+">", str)
		}
		select {
		case <-r.notifyChanFinish:
			//一次性读取所有剩余数据
			fmt.Print("<"+r.id+">", r.outBuf.String())

			RunnerAll.Del(r.id)
			return
		default:
			continue
		}
	}
}
