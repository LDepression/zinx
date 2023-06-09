/**
 * @Author: lenovo
 * @Description:
 * @File:  connection
 * @Version: 1.0.0
 * @Date: 2023/04/11 14:33
 */

package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

type Connection struct {
	//关联的Serve
	Serve ziface.IServer
	//当前连接的socket TCP套接字
	Conn *net.TCPConn
	//连接的ID
	ConnID uint32

	//当前连接的状态
	isClosed bool

	//无缓冲通道，用于读，写Goroutine之间通信
	msgChan chan []byte
	//告知当前连接已经退出的的channel
	ExitChan  chan bool
	HandleMsg ziface.IMsgHandler
}

//初始化连接模块的方法

func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandler) *Connection {
	c := &Connection{
		Serve:     server,
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		HandleMsg: msgHandler,
		msgChan:   make(chan []byte),
		ExitChan:  make(chan bool, 1),
	}
	c.Serve.GetConnMgr().Add(c)
	return c
}

// 连接中读取业务的方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println("ConnID =", c.ConnID, "Reader is exit ,remote addr is ", c.RemoteAddr())
	defer c.Stop()

	for {
		//读取客户端的数据到buf中
		//buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		//_, err := c.Conn.Read(buf) //读到buf中去
		//if err != nil {
		//	fmt.Println("recv  buf err", err)
		//	continue
		//}

		//创建一个拆包/解包的对象
		dp := NewDataPack()

		//读取客户端的Msg Head 二级字节流 8个字节
		headData := make([]byte, dp.GetHeadLen())
		//根据dataLen,再次读取Data,放在Msg.Data中去
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head err:", err)
			break
		}

		//拆包,得到MsgID和MsgDataLe放在Msg消息中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("read msg head error,err:", err)
			break
		}

		//根据dataLen,再次读取Data,放在msg.Data中去
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data err:", err)
				break
			}
		}
		msg.SetData(data)
		//根据当前conn数据的Request请求数据

		//从路由中,找到注册绑定的Conn对应的Router调用
		req := Request{
			conn: c,
			msg:  msg,
		}
		if utils.GlobalObject.WorkerPoolSize > 0 {
			c.HandleMsg.SendMsgToQueue(&req)
		} else {
			//根据绑定好的MsgID找到处理好的api业务 执行
			go c.HandleMsg.DoMsgHandler(&req)
		}
	}
}

/*
写消息Goroutine，用户将消息发送给客户端
*/
func (c *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is running] ")
	defer fmt.Println(c.RemoteAddr().String(), "conn has Exited")

	//不断阻塞，等待chan的消息，从而发给客户端

	for {
		select {
		case data := <-c.msgChan:
			//有数据的话，就要写给客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("send Msg err", err)
				return
			}
		case <-c.ExitChan:
			//代表reader已经退出了，此时Writer也要退出
			return
		}
	}
}

// 提供一个sendMsg方法,将我们要给客户端的数据先封包,再去发送]
func (c *Connection) SendMsg(MsgID uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("connection closed when send msg")
	}
	//将data进行封包 MsgDataLen|MsgID\MSG
	dp := NewDataPack()

	binaryMsg, err := dp.Pack(NewMsgPackage(MsgID, data))
	if err != nil {
		fmt.Println("pack msg id =", MsgID)
		return errors.New("pack msg err")
	}

	//将数据发送给客户端
	c.msgChan <- binaryMsg
	return nil
}

// 启动连接,让当前的连接准备开始工作
func (c *Connection) Start() {
	fmt.Println("Connect Start .... ConnID=", c.ConnID)

	//启动从当前连接中读取业务
	go c.StartReader()

	//todo:写业务
	go c.StartWriter()
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
	c.ExitChan <- true
	c.Conn.Close()

	//将当前的Conn从Mgr中清理掉
	c.Serve.GetConnMgr().Remove(c)
	close(c.ExitChan)
	close(c.msgChan)
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
