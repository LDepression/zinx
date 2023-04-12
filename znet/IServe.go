/**
 * @Author: lenovo
 * @Description:
 * @File:  IServe
 * @Version: 1.0.0
 * @Date: 2023/04/11 13:25
 */

package znet

import (
	"fmt"
	"net"
	"zinx/utils"
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

	//当前的Server添加一个router
	Router ziface.IRouter
}

// 启动服务器的方法
func (s *Server) Start() {
	fmt.Printf("[ZINX] Server Name :%s, listenner at IP: :%s, port :%d", utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
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
			dealConn := NewConnection(conn, cid, s.Router)
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

// 添加一个路由
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add Router Succ!!")
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:    nil,
	}
	return s
}
