/**
 * @Author: lenovo
 * @Description:
 * @File:  idatapack
 * @Version: 1.0.0
 * @Date: 2023/04/12 15:53
 */

package ziface

/*
封包,拆包 模块
用于处理TCP黏包问题
*/
type IDataPack interface {
	//获取包头的长度的方法
	GetHeadLen() uint32
	//封包方法
	Pack(msg IMessage) ([]byte, error)
	//拆包方法
	Unpack([]byte) (IMessage, error)
}
