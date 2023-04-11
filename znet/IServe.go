/**
 * @Author: lenovo
 * @Description:
 * @File:  IServe
 * @Version: 1.0.0
 * @Date: 2023/04/11 13:25
 */

package znet

import (
	"errors"
	"fmt"
	"net"
	"zinx/ziface"
)

//iServe的接口实现

type Server struct {
	//服务器的名称
	Name string
	//服务器绑定的ip版本
	IPVersion string
	//服务器监听的ip
	IP string
	//服务器监听的端口
	Port int
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	//回显的业务
	fmt.Println("[Conn Handle] CallBackToClient...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err", err)
		return errors.New("CallBackToClient error")
	}
	return nil
}

// 启动服务器的方法
func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at Ip :%s , Port is %d, is starting \n", s.IP, s.Port)

	go func() {
		//1.获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolveTCPAddr error: ", err)
			return
		}
		//2.监听服务器的地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, "err ", err)
			return
		}
		fmt.Println("start Zinx server succ,", s.Name, "succ Listening")
		var cid uint32 = 0

		//3.阻塞等待客户端连接,处理客户端的业务
		for {
			//如果有客户端连接过来,阻塞会返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept,err", err)
				continue
			}

			//将处理新链接的业务方法和Conn进行绑定,得到我们得到连接模块
			dealConn := NewConnection(conn, cid, CallBackToClient)
			cid++
			//启动当前连接业务处理
			go dealConn.Start()
		}
	}()
}

// 停止服务器的方法
func (s *Server) Stop() {
	//todo: 将一些服务器的资源,状态或者一些已经开辟的连接信息进行停止或者回收

}

// 运行服务器的方法
func (s *Server) Serve() {
	s.Start()

	//todo: 做一些额外的业务
	//阻塞
	select {}

}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
