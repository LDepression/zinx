/**
 * @Author: lenovo
 * @Description:
 * @File:  msgHandler
 * @Version: 1.0.0
 * @Date: 2023/04/13 21:23
 */

package ziface

type IMsgHandler interface {

	//调度/执行对应router消息处理方法
	DoMsgHandler(request IRequest)
	//为消息添加具体的处理逻辑
	AddRouter(MsgID uint32, router IRouter)
}
