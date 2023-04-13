/**
 * @Author: lenovo
 * @Description:
 * @File:  msgHandler
 * @Version: 1.0.0
 * @Date: 2023/04/13 21:22
 */

package znet

import (
	"fmt"
	"zinx/ziface"
)

type MsgHandler struct {
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandle() *MsgHandler {
	return &MsgHandler{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

// 调度/执行对应router消息处理方法
func (mh *MsgHandler) DoMsgHandler(request ziface.IRequest) {
	//找到当前收到请求的msgID
	handler, ok := mh.Apis[request.GetMsgID()]
	if ok == false {
		fmt.Println("api msgID =", request.GetMsgID(), "not found Need Register")
	}

	//有的话,直接调用
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

// 为消息添加具体的处理逻辑
func (mh *MsgHandler) AddRouter(MsgID uint32, router ziface.IRouter) {

	//判断一下当前的msg绑定的ID是否已经存在了
	if _, ok := mh.Apis[MsgID]; ok {
		panic("repeat message ID")
	}
	//添加Msg与api的关系
	mh.Apis[MsgID] = router
	fmt.Println("Add MsgID", MsgID, "succ")

}
