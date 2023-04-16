/**
 * @Author: lenovo
 * @Description:
 * @File:  connectionmanager
 * @Version: 1.0.0
 * @Date: 2023/04/16 16:30
 */

package znet

import (
	"errors"
	"fmt"
	"sync"
	"zinx/ziface"
)

/*
连接管理模块
*/
type ConnectionManager struct {
	Connections map[uint32]ziface.IConnection
	ConnLock    sync.RWMutex
}

func NewConnManager() *ConnectionManager {
	return &ConnectionManager{
		Connections: make(map[uint32]ziface.IConnection),
		ConnLock:    sync.RWMutex{},
	}
}

// 添加连接
func (connMgr *ConnectionManager) Add(conn ziface.IConnection) {
	connMgr.ConnLock.Lock()
	defer connMgr.ConnLock.Unlock()

	connMgr.Connections[conn.GetConnID()] = conn
	fmt.Println("Conn add to ConnManager succ,connNum:", connMgr.Len())
}

// 删除连接
func (connMgr *ConnectionManager) Remove(conn ziface.IConnection) {
	connMgr.ConnLock.RLock()
	defer connMgr.ConnLock.RUnlock()
	delete(connMgr.Connections, conn.GetConnID())
}

// 根据ConnID获取连接
func (connMgr *ConnectionManager) Get(connID uint32) (ziface.IConnection, error) {
	if conn, ok := connMgr.Connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("get connection failed")
	}
}

// 得到当前的连接个数
func (connMgr *ConnectionManager) Len() uint32 {
	return uint32(len(connMgr.Connections))
}

// 清除并终止所有连接
func (connMgr *ConnectionManager) ClearConn() {
	for connID, conn := range connMgr.Connections {
		conn.Stop()
		delete(connMgr.Connections, connID)
	}
	fmt.Println("clear all connections, now connNum=", connMgr.Len())
}
