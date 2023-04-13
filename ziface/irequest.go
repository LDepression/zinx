/**
 * @Author: lenovo
 * @Description:
 * @File:  irequest
 * @Version: 1.0.0
 * @Date: 2023/04/12 12:31
 */

package ziface

type IRequest interface {

	//得到当前的连接
	GetConnection() IConnection
	//得到请求的消息数据
	GetData() []byte

	//得到请求的ID
	GetMsgID() uint32
}
