/**
 * @Author: lenovo
 * @Description:
 * @File:  imessage
 * @Version: 1.0.0
 * @Date: 2023/04/12 15:27
 */

package ziface

/*

将请求的消息封装到一个Message中,
定义抽象的接口
*/

type IMessage interface {
	GetMsgID() uint32  //获取消息的ID
	GetMsgLen() uint32 //获取消息的长度
	GetData() []byte   //获取消息的内容
	SetData([]byte)    //设置消息的内容
	SetDataLen(uint32) //设置消息的长度
}
