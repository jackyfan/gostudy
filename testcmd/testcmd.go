// testcmd
package testcmd

import (
	"flag"
	"fmt"
)

func Testcmd() {
	//第一个参数是“参数名”，第二个是“默认值”，第三个是“说明”。返回的是指针
	host := flag.String("host", "coolshell.cn", "a host name ")
	port := flag.Int("port", 80, "a port number")
	debug := flag.Bool("d", false, "enable/disable debug mode")

	//正式开始Parse命令行参数
	flag.Parse()

	fmt.Println("host:", *host)
	fmt.Println("port:", *port)
	fmt.Println("debug:", *debug)
}
