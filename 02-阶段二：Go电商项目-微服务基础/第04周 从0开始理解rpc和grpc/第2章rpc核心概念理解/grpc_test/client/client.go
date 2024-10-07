/**
 * @File : client.go
 * @Description : 请填写文件描述
 * @Author : 请填写作者的真实姓名
 * @Date : 2024-10-06
 */
package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		panic("connect fail")
	}

	var reply string
	err = client.Call("HelloService.Hello", "world", &reply)
	if err != nil {
		panic("call fail")
	}
	fmt.Println(reply)
}
