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
	"zinx/utils"
	"zinx/ziface"
)

type MsgHandler struct {
	Apis           map[uint32]ziface.IRouter
	TaskQueue      []chan ziface.IRequest
	WorkerPoolSize uint32
}

func NewMsgHandle() *MsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
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

// 启动一个worker工作池,只能启动一次
func (mh *MsgHandler) StartWorkerPool() {
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		go mh.StartOneWork(i, mh.TaskQueue[i])
	}
}

// 启动工作流程
func (mh *MsgHandler) StartOneWork(workerID int, tackQueue chan ziface.IRequest) {
	fmt.Println("workerID =", workerID, "is started")
	for {
		select {
		case request := <-tackQueue:
			mh.DoMsgHandler(request)
		}
	}

}

func (mh *MsgHandler) SendMsgToQueue(request ziface.IRequest) {
	//1.将消息平均分配给不同的Worker
	WorkerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	//将消息塞进消息队列中
	fmt.Println("Add ConnID=", request.GetConnection().GetConnID(),
		"request msgID", request.GetMsgID(),
		"to WorkerID", WorkerID,
	)

	mh.TaskQueue[WorkerID] <- request
}
