package znet

import (
	"fmt"
	"my-zinx/ziface"
	"net"
)

type Connection struct {
	// 当前链接的 socket TCP 套接字
	Conn *net.TCPConn

	// 链接的 ID
	ConnId uint32

	// 当前链接状态
	isClosed bool

	// 当前链接所绑定的处理业务的方法 API
	handleAPI ziface.HandleFunc

	// 告知当前链接已经停止/推出的 channel
	ExitChan chan bool
}

// NewConnection 初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connId uint32, callbackApi ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnId:    connId,
		handleAPI: callbackApi,
		isClosed:  false,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

// StartReader 链接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")

	defer fmt.Println("ConnId = ", c.ConnId, "Reader is exit, remote add is", c.RemoteAddr().String())
	defer c.Stop()

	for {
		// 读取客户端的数据到 buf 中，目前最大是 512 字节
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("Receive buf err", err)
			break
		}

		// 调用当前链接所绑定的 HandleAPI
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnId", c.ConnId, "handle is error", err)
			break
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("Conn Start().. ConnId = ", c.ConnId)

	// 启动从当前链接的读数据的业务
	go c.StartReader()

	// TODO 启动从当前链接写数据的业务
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop().. ConnId = ", c.ConnId)

	// 如果当前链接已经关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	// 调用关闭 socket 链接
	err := c.Conn.Close()
	if err != nil {
		panic("Conn Stop error!")
	}

	// 回收资源
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	//TODO implement me
	panic("implement me")
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnId
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	//TODO implement me
	panic("implement me")
}
