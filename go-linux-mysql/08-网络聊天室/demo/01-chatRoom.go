package main

import (
	"net"
	"fmt"
	"io"
	"strings"
	"time"
)

// 创建用户结构体
type Client struct {
	Name string
	Addr string
	C chan string
}

// 创建全局map，没有空间
var onLineMap map[string]Client

// 创建全局 channel
var message = make(chan string)

func MakeMsg(client Client, str string) string {
	msg := "[" + client.Addr + "]" + client.Name + ": " + str
	return msg
}

func HandlerConnect(conn net.Conn)  {
	defer conn.Close()
	fmt.Println("--------------------------\n")
	// 组织用户信息，获取用户的IP+port. 初始Name == Addr
	netAddr := conn.RemoteAddr().String()
	clit := Client{netAddr, netAddr, make(chan string)}

	// 将用户添加到 在线用户列表中 onLineMap
	onLineMap[netAddr] = clit

	// 创建 读取用户自带channel go程
	go WriteMsgToClient(conn, clit)

	// 组织用户上线消息
	//msg := "[" + netAddr + "]" + clit.Name + ":login\n"
	msg := MakeMsg(clit, "login")

	// 写 用户上线消息 到 全局channel
	message <- msg

	// 创建一个channel 协调go程退出顺序。
	isQuit := make(chan bool)
	// 创建一个证实用户是 活跃用户的channel
	isLife := make(chan bool)

	// 创建 监听用户聊天信息 go 程 —— 匿名go程
	go func() {
		buf := make([]byte, 4096)
		// 循环 读取conn上客户端发送的数据。—— 无：阻塞等待
		for {
			n, err := conn.Read(buf)
			if n == 0 { // client 端 close()
				fmt.Println("客户端下线")
				isQuit <- true 			// 通知 HandlerConnet go程 结束。
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Read err:", err)
				return
			}
			// 判断用户发送的是否是一个 查询在线用户命令
			length := len(buf[:n-1])
			if string(buf[:n-1]) == "who" && length == 3 {
				conn.Write([]byte("OnLine User list:\n"))
				for _, client := range onLineMap {
					// 将 查询到的在线用户写回给当前对应客户端。 conn.write
					msg := MakeMsg(client, " [OnLine]\n")
					conn.Write([]byte(msg))
				}
				// 判断用户 发送的是一个 改名 命令  ：  rename|新用户名
			} else if (string(buf[:7]) == "rename|" && length >= 8) {
				// 提取用户指定的新用户名, split 拆分
				newName := strings.Split(string(buf[:n-1]), "|")[1]
				// 将新用户名替换旧用户名
				clit.Name = newName
				// 将改名后的用户结构体，替换掉 onLineMap 中该用户源结构体。
				onLineMap[netAddr] = clit
				// 给用户提示，改名成功
				conn.Write([]byte("rename successful!!!\n"))
			}else {
				// 成功读到客户端数据。写数据到全局channel中 —— 广播
				msg := MakeMsg(clit, string(buf[:n]))
				// 写 用户聊天消息 到 全局channel
				message <- msg
			}
			isLife <- true		// 表示当前用户为活跃用户。
		}
	}()
/*	// 防止当前 HandlerConnect 退出
	for {
		runtime.GC();
	}*/

	// 循环 select 判断是否用户退出
	for {
		select {
			case <-isQuit: // 客户端close。
				// 从在线用户列表中踢出当前用户。
				delete(onLineMap, clit.Addr)
				//delete(onLineMap, netAddr)

				// 广播给所有在线用户，当前用户下线。
				msg := MakeMsg(clit, "logout!")
				message <- msg
				//break   // 不能使用 break
				// 关闭用户自带channel 写端。迫使WriteMsgToClient go程 结束range循环，主动退出
				close(clit.C)
				return
			case <-isLife:
				// 什么也不做！ 目的就是为了重置下面的计时器
			case <-time.After(time.Second * 60 * 60 * 60 * 2):
				delete(onLineMap, clit.Addr)
				// 广播给所有在线用户，当前用户超时退出 被迫下线。
				msg := MakeMsg(clit, "time out to leave !")
				message <- msg
				close(clit.C)
				return
		}
	}
}

// 读取用户自带 channel
func WriteMsgToClient(conn net.Conn, clit Client)  {
	// 循环读取用户自带 channel 上是否有数据
/*	for {
		msg := <-clit.C		// 无——阻塞。 有——继续
		conn.Write([]byte(msg + "\n"))
	}*/
	for msg := range clit.C {				// 判断 用户自带channel中是否有数据可读。
		conn.Write([]byte(msg + "\n"))
	}
}

// 在线用户管理go程
func Manager()  {
	// 初始化 在线用户列表空间
	onLineMap = make(map[string]Client)

	// 循环 监听全局channel上，是否有数据，如果没有阻塞，如果有读出，写给每一个用户。、
	for {
		msg := <-message

		// 遍历map，写出给每一个用户
		for _, clit := range onLineMap {
			clit.C <- msg			// 将从全局channel中读到的数据，写给当前用户。
		}
	}
}

func main()  {
	// 创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8009")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	// 创建 在线用户管理go程
	go Manager()

	// 循环阻塞等待客户端连机请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err:", err)
			continue
		}
		// 创建处理用户消息及相关动作的 go程
		go HandlerConnect(conn)
	}
}
