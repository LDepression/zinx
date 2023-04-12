/**
 * @Author: lenovo
 * @Description:
 * @File:  client
 * @Version: 1.0.0
 * @Date: 2023/04/11 13:58
 */

package main

import (
	"bytes"
	"fmt"
	"net"
	"time"
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
		//2.连接调用Write方法
		_, err := conn.Write([]byte("Hello, ZinxV1.0!"))
		if err != nil {
			fmt.Println("write conn err", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}
		buf = bytes.TrimSpace(buf[:cnt])
		fmt.Println("server callback: [", string(buf), "] cnt: [", cnt, "]")

		//cpu阻塞
		time.Sleep(time.Second)
	}
}
