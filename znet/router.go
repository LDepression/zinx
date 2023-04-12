/**
 * @Author: lenovo
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2023/04/12 12:38
 */

package znet

import "zinx/ziface"

// 实现router时,先嵌入这个Base基类,然后根据需要对这个基类的方法重写就好了

/*
之所以BaseRouter方法为空,
因为有的Router不希望有PreHandle ,postHandle这两个业务
Router全部继承BaseRouter的好处是,不需要实现PreHandle
*/

type BaseRouter struct {
}

func (br *BaseRouter) PreHandle(request ziface.IRequest) {

}

// 在处理conn业务的主方法
func (br *BaseRouter) Handle(request ziface.IRequest) {

}

// 在处理conn业务知乎的钩子方法Hook
func (br *BaseRouter) PostHandle(request ziface.IRequest) {

}
