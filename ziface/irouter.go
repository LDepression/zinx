/**
 * @Author: lenovo
 * @Description:
 * @File:  irouter
 * @Version: 1.0.0
 * @Date: 2023/04/12 12:38
 */

package ziface

/*
路由的抽象接口,
路由里面的数据IRequest
*/
type IRouter interface {
	//在处理conn业务之前的钩子方法Hook
	PreHandle(request IRequest)
	//在处理conn业务的主方法
	Handle(request IRequest)
	//在处理conn业务知乎的钩子方法Hook
	PostHandle(request IRequest)
}
