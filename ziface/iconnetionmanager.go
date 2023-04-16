/**
 * @Author: lenovo
 * @Description:
 * @File:  iconnetionmanager
 * @Version: 1.0.0
 * @Date: 2023/04/16 16:30
 */

package ziface

type IConnManager interface {
	//添加连接
	Add(conn IConnection)
	//删除连接
	Remove(conn IConnection)
	//根据ConnID获取连接
	Get(connID uint32) (IConnection, error)
	//得到当前的连接个数
	Len() uint32
	//清除并终止所有连接
	ClearConn()
}
