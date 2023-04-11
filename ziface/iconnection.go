/**
 * @Author: lenovo
 * @Description:
 * @File:  iconnection
 * @Version: 1.0.0
 * @Date: 2023/04/11 14:33
 */

package ziface

import "net"

// 定义抽象模块的抽象层
type IConnection interface {

	//启动连接,让当前的连接准备开始工作
	Start()
	//停止工作
	Stop()
	//获取当前连接绑定的socket conn
	GetTCPConnection() *net.TCPConn
	//获取当前连接模块的连接ID
	GetConnID() uint32
	//获取远程客户端的TCP的IP port
	RemoteAddr()

	//发送数据
	Send(data []byte) error
}

// 定义一个处理连接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
