/**
 * @Author: lenovo
 * @Description:
 * @File:  IServe
 * @Version: 1.0.0
 * @Date: 2023/04/11 13:25
 */

package ziface

//定义一个服务器接口

type IServer interface {
	Start()
	Stop()
	Serve()

	//路由功能:给当前服务注册一个理由方法,供客户端使用
	AddRouter(msgID uint32, router IRouter)
	//得到连接管理器
	GetConnMgr() IConnManager
}
