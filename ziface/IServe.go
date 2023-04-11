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
}
