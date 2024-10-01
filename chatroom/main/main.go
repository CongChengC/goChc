package main
import (
	"fmt"
	"net"
	//"goChc/process"
	//"goChc/chatroom/server/service"
	"goChc/chatroom/server/service"
)

func main() {

	// 假设 conn 是一个已经打开的 net.Conn
	fmt.Println("服务器[新的结构]在8889端口监听....")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()

	if err != nil {
		fmt.Println("net.Listen err=", err)
		return 
	}

	conn, err := listen.Accept()

	// pro := &process.Pro{
	// 	Conn : conn,
	// }

	pro := &service.Pro2{
		Conn : conn,
	}

	//提示信息
	fmt.Println("服务器[新的结构]在8889端口监听....")
	fmt.Println("服务器[新的结构]在8889端口监听....")

	fmt.Println(pro)
	
}