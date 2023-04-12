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

func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router Prehandle")
	request.GetConnection().GetTCPConnection().Write([]byte("Before ping .......\n"))
}
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle")
	request.GetConnection().GetTCPConnection().Write([]byte(" ping ...ping ...ping...\n"))

}
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router Prehandle")
	request.GetConnection().GetTCPConnection().Write([]byte("after ping .......\n"))
}
func main() {
	//1.创建一个Serve句柄,使用Zinx的API
	s := znet.NewServer("[ZinxV0.3]")

	//添加Router
	s.AddRouter(&PingRouter{})
	//3.启动serve
	s.Serve()
}
