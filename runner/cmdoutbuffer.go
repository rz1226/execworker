package runner

import (
	"bytes"

	"sync"
)

//和bytes.Buffer的区别是对用到的几个方法增加mutex。因为cmd执行的时候往里写，同时要从里面取数据作为日志

const (
	CMDOUTBUFFERSIZE = 1000
)

//带锁的
type CmdOutBuffer struct {
	buf *bytes.Buffer
	mu  *sync.Mutex
}

func newCmdOutBuffer() *CmdOutBuffer {
	m := &CmdOutBuffer{}
	m.buf = bytes.NewBuffer(make([]byte, 0, CMDOUTBUFFERSIZE))
	m.mu = &sync.Mutex{}
	return m
}

func (m *CmdOutBuffer) ReadString(delim byte) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	line, err := m.buf.ReadString(delim)
	return line, err
}

func (m *CmdOutBuffer) Write(p []byte) (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	n, err := m.buf.Write(p)
	return n, err
}

func (m *CmdOutBuffer) ReadRune() (rune, int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	r, size, err := m.buf.ReadRune()
	return r, size, err
}

func (m *CmdOutBuffer) String() string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.buf.String()

}
