/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/04/11 13:33
 */

package main

import "zinx/znet"

/*
	基于Zinx框架来开发 服务器端应用程序

*/

func main() {
	//1.创建一个Serve句柄,使用Zinx的API
	s := znet.NewServer("[ZinxV0.1]")
	//2.启动serve
	s.Serve()
}
