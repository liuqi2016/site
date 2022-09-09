package site

import (
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestSign(t *testing.T) {
	sigRecv := make(chan os.Signal, 1)
	signal.Notify(sigRecv)
	for sig := range sigRecv {
		t.Log(sig)
	}
}

func SendSign(t testing.T) {
	//首先找到当前进程
	pid := 10086
	ps, _ := os.FindProcess(pid)
	ps.Signal(syscall.SIGINT)
}
