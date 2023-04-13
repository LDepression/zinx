/**
 * @Author: lenovo
 * @Description:
 * @File:  client
 * @Version: 1.0.0
 * @Date: 2023/04/11 13:58
 */

package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/znet"
)

/*

模拟客户端
*/

func main() {
	fmt.Println("client start....")
	time.Sleep(1 * time.Second)
	//1.直接连接远程服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err,exit")
		return
	}

	for {
		//发送封包的Msg消息
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMsgPackage(0, []byte("ZinxV0.5 Client Test Request")))
		if err != nil {
			fmt.Println("Pack err:", err)
			return
		}
		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("Write err:", err)
			return
		}

		//服务器应该就应该给我们回复一个数据msgID 为1 的ping...ping...ping...
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head err:", err)
			return
		}

		//将二进制的head拆封到结构体去
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("client unpacked msg:", err)
			return
		}
		if msgHead.GetMsgLen() > 0 {
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg err:", err)
				return
			}
			fmt.Println("---> Recv Server Msg : ID = ", msg.ID, ", len = ", msg.DataLen, ", data = ", string(msg.Data))
		}
		//cpu阻塞
		time.Sleep(time.Second)
	}
}
