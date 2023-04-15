/**
 * @Author: lenovo
 * @Description:
 * @File:  globalobj
 * @Version: 1.0.0
 * @Date: 2023/04/12 14:06
 */

package utils

import (
	"encoding/json"
	"os"
	"zinx/ziface"
)

type GlobalObj struct {

	/*
		Server
	*/
	TcpServer ziface.IServer //当前Zinx全局的Server对象
	Host      string         //当前服务器监听的IP
	TcpPort   int            //当前服务器监听的端口

	Name string //当前服务器的名称

	/*
		Zinx'
	*/
	Version          string //当前Zinx的版本号
	MaxConn          uint32 //当前服务器最大连接数
	MaxPackageSize   uint32
	WorkerPoolSize   uint32 //当前业务工作池goroutine的数量
	MaxWorkerTaskLen uint32 // 每个worker消息队列最大的长度
}

var GlobalObject *GlobalObj

func (g GlobalObj) Reload() {
	data, err := os.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, GlobalObject); err != nil {
		panic(err)
	}
}
func init() {
	//如果配置文件没有加载，默认的值
	GlobalObject = &GlobalObj{
		Name:             "ZinxServerApp",
		Version:          "V0.6",
		TcpPort:          8999,
		Host:             "0.0.0.0",
		MaxConn:          1000,
		MaxPackageSize:   4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
	}

	//应该尝试从conf/zinx.json去加载一些用户自定义的参数
	GlobalObject.Reload()
}
