package main

import "my-zinx/znet"

func main() {
	// 创建一个 Server 的句柄，使用 Zinx 的 api
	s := znet.NewServer("[zinx v0.1]")

	// 启动 Server
	s.Serve()
}
