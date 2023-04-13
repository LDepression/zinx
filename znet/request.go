/**
 * @Author: lenovo
 * @Description:
 * @File:  request
 * @Version: 1.0.0
 * @Date: 2023/04/12 12:31
 */

package znet

import "zinx/ziface"

type Request struct {
	//已经和客户端建立好的连接
	conn ziface.IConnection
	//客户端请求的数据
	msg ziface.IMessage
}

// 得到当前的连接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}
