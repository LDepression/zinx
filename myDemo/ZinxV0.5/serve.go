/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/04/11 13:33
 */

package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

/*
	基于Zinx框架来开发 服务器端应用程序

*/

//ping test

type PingRouter struct {
	znet.BaseRouter
}

//Test PreHandle

func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle")
	//先读取客户端的数据,再写回ping...ping...ping...
	fmt.Println("recv from client, MsgID: ", request.GetMsgID(), "data:", string(request.GetData()))
	if err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping")); err != nil {
		fmt.Println(err)
	}
}

func main() {
	//1.创建一个Serve句柄,使用Zinx的API
	s := znet.NewServer("[ZinxV0.5]")

	//添加Router
	s.AddRouter(&PingRouter{})
	//3.启动serve
	s.Serve()
}
