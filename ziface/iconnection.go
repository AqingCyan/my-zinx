package ziface

import "net"

// IConnection 定义链接模块的抽象层
type IConnection interface {
	// Start 启动链接，让当前的链接准备开始工作
	Start()

	// Stop 停止链接，结束当前链接的工作
	Stop()

	// GetTCPConnection 获取当前链接所绑定的 socket connection
	GetTCPConnection() *net.TCPConn

	// GetConnID 获取当前链接模块的链接 ID
	GetConnID() uint32

	// RemoteAddr 获取远程客户端的 TCP 状态以及 IP port
	RemoteAddr() net.Addr

	// Send 发送数据，将数据发送给远程的客户端
	Send(data []byte) error
}

// HandleFunc 定义一个处理链接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
