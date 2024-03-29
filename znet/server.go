package znet

import (
	"fmt"
	"my-zinx/utils"
	"my-zinx/ziface"
	"net"
)

// Server iServer 的接口实现，定义一个 Server 的服务器模块
type Server struct {
	// 服务器名称
	Name string

	// 服务器绑定的 IP 版本
	IPVersion string

	// 服务器监听的 IP
	IP string

	// 服务器监听的端口
	Port int

	// 当前的 Server 添加一个 router，server 注册的链接对应的处理业务
	Router ziface.IRouter
}

// NewServer 初始化 Server 模块的方法
func NewServer() ziface.IServer {
	return &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:    nil,
	}
}

func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Name: %s, listener at IP: %s, Prot: %d is starting\n", utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[Zinx] Version %s, MaxConn: %d, MaxPackageSize: %d\n", utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPackageSize)

	go func() {
		// 1 获取一个 TCP 的 Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("Resolve tcp addr error", err)
			return
		}

		// 2 监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("Listen", s.IPVersion, "err", err)
			return
		}

		fmt.Println("Start Zinx server success, ", s.Name, "success, Listening...")
		var cid uint32
		cid = 0

		// 3 阻塞的等待客户端连接，处理客户端链接的业务（读写）
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			// 将处理新链接的业务方法和 conn 进行绑定，得到我们的链接模块
			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			// 启动当前的链接业务处理
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Serve() {
	// 启动 server 的服务功能
	s.Start()

	// TODO 做一些启动服务器之后的额外业务

	// 阻塞住，因为服务启动在 go func 里
	select {}
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add Router Success!!")
}
