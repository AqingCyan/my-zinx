package utils

import (
	"encoding/json"
	"my-zinx/ziface"
	"os"
)

// GlobalObj 存储一切有关 Zinx 框架的全局参数，供其他模块使用  一些参数是可以通过 zinx.json 由用户进行配置
type GlobalObj struct {
	// Server 部分
	TcpServer ziface.IServer // 当前 Zinx 全局的 Server 对象
	Host      string         // 当前服务主机监听的 IP
	TcpPort   int            // 当前服务器主机监听的端口号
	Name      string         // 当前服务器的名称

	// Zinx 部分
	Version        string // 当前 Zinx 的版本号
	MaxConn        int    // 当前服务器主机允许的最大链接数
	MaxPackageSize uint32 // 当前 Zinx 框架数据包的最大值
}

// GlobalObject 定义一个全局的对外的 GlobalObj
var GlobalObject *GlobalObj

// Reload 从 zinx.json 去加载用于自定义的参数
func (g *GlobalObj) Reload() {
	data, err := os.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}

	// 将 json 文件数据解析到 struct 中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

// init 提供一个方法初始化当前 GlobalObject
func init() {
	// 没有配置文件就默认加载的值
	GlobalObject = &GlobalObj{
		Name:           "ZinxServerApp",
		Version:        "0.4",
		TcpPort:        8889,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	// 应该尝试从 conf/zinx.json 去加载一些用户自定义的配置
	GlobalObject.Reload()
}
