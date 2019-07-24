package runner

import (
	"fmt"
	"sync"
)

//存放所有正在进行的任务
var RunnerAll *RunnerList

func init(){
	RunnerAll = NewRunnerList(300)
}

func Info() string {
	return RunnerAll.Info()
}





/*******--------------------***/
//所有runner实例的集合
type RunnerList struct{
	m   sync.Mutex
	rs    map[string]*Runner
}
func NewRunnerList(size int )*RunnerList{
	rl := &RunnerList{}
	rl.rs = make(map[string]*Runner, size)
	//rl.m = new(sync.Mutex)
	return rl
}

func (rl *RunnerList) Add( r *Runner){
	rl.m.Lock()
	defer rl.m.Unlock()
	rl.rs[r.id] = r
}

func (rl *RunnerList) Del( id string){
	rl.m.Lock()
	defer rl.m.Unlock()
	delete( rl.rs, id )
}

func (rl *RunnerList)  Info() string {
	rl.m.Lock()
	defer rl.m.Unlock()
	str := ""
	for _, v := range rl.rs {

		str += fmt.Sprintln("start time:", v.startTime)
		str += fmt.Sprintln("command:",  v.cmd.Args)
	}
	return str
}

