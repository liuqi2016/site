---
title: "golang并发编程实战一" # Title of the blog post.
date: 2022-09-07T22:29:45+08:00 # Date of post creation.
description: "golang并发编程实阅读与总结,精简出高质量的内容." # Description used for search engine.
tags:
  - golang
  - 并发
---
# pipe 管道
父进程与子进程或者同祖进程间通信,单向  
io.pipe(系统级别的管道,不支持原子操作)、os.pipe(内存管道、使用了sync)
wrire和read执行数据时会等待另一端准备就绪,需要并发执行,否则会阻塞报死锁。


# signal 信号
唯一一个异步通信，软件模拟硬件中断，事件通知进程。
kill-l 一共62种 
1-31可靠信号
34-64不可靠信号
![sign](/static/images/sign.png)


```golang
//自处理系统信号,Notify第二个参数传信号类型
//案例屏蔽所有信号,除了sigkill和sigstop不能屏蔽
func TestSign(t *testing.T) {
	sigRecv := make(chan os.Signal, 1)
	signal.Notify(sigRecv)
	for sig := range sigRecv {
		t.Log(sig)
	}
}

//使用程序向进程发送信号
func SendSign(t testing.T) {
	//首先找到当前进程
	pid := 10086
	ps, _ := os.FindProcess(pid)
	ps.Signal(syscall.SIGINT)
}

```
# socket 嵌套字
通过网络连接让多个进程通信


