/**
 * @Author: lenovo
 * @Description:
 * @File:  connection
 * @Version: 1.0.0
 * @Date: 2023/04/11 14:33
 */

package znet

import (
	"fmt"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

type Connection struct {
	//当前连接的socket TCP套接字
	Conn *net.TCPConn
	//连接的ID
	ConnID uint32

	//当前连接的状态
	isClosed bool

	//告知当前连接已经退出的的channel
	ExitChan chan bool

	//该链接处理的方法router
	Router ziface.IRouter
}

//初始化连接模块的方法

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		Router:   router,
		ExitChan: make(chan bool, 1),
	}
	return c
}

// 连接中读取业务的方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println("ConnID =", c.ConnID, "Reader is exit ,remote addr is ", c.RemoteAddr())
	defer c.Stop()

	for {
		//读取客户端的数据到buf中
		buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		_, err := c.Conn.Read(buf) //读到buf中去
		if err != nil {
			fmt.Println("recv  buf err", err)
			continue
		}

		//从路由中,找到注册绑定的Conn对应的Router调用
		req := Request{
			conn: c,
			data: buf,
		}
		//执行注册路由的方法
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

// 启动连接,让当前的连接准备开始工作
func (c *Connection) Start() {
	fmt.Println("Connect Start .... ConnID=", c.ConnID)

	//启动从当前连接中读取业务
	go c.StartReader()

	//todo:写业务
}

// 停止工作
func (c *Connection) Stop() {
	fmt.Println("Conn Stop() ConnID:", c.ConnID)
	//如果当前连接已经关闭的话
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	//挂麻痹socket连接
	c.Conn.Close()
	close(c.ExitChan)
}

// 获取当前连接绑定的socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前连接模块的连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端的TCP的IP port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据
func (c *Connection) Send(data []byte) error {
	return nil
}
