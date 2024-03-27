package main

import (
	"fmt"
	"my-zinx/ziface"
	"my-zinx/znet"
)

// PingRouter 做一个 ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// PreHandle Test
func (pr *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("Call back ping error")
	}
}

// Handle Test
func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping...\n"))
	if err != nil {
		fmt.Println("Call back ping...ping...ping... error")
	}
}

// PostHandle Test
func (pr *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("after ping...")
	}
}

func main() {
	// 创建一个 Server 的句柄，使用 Zinx 的 api
	s := znet.NewServer("[zinx v0.3]")

	// 给当前 zinx 添加一个自定义的 router
	s.AddRouter(&PingRouter{})

	// 启动 Server
	s.Serve()
}
