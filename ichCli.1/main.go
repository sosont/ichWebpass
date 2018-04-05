package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var (
	svrAddr = flag.String("addr", "127.0.0.1", "为连接服务器的地址")

	tcpPort        = flag.Int("tp", 0, "Socket连接server的端口")
	remoteHttpProt = flag.Int("rhp", 0, "穿透远程开启HTTP端口")
	httpPort       = flag.Int("hp", 0, "为转发至本地客户端的端口")
	verifyKey      = flag.String("vkey", "", "用作客户端与服务端连接时的校验")
)

func main() {
	fmt.Println("+----------------------------------------------------------------+")
	fmt.Println("| Welcome to ichWebpass Ver1.0.0 .                               |")
	fmt.Println("| Code by yaoqi at 2018-04-06                                    |")
	fmt.Println("| If you have some problem when you use the tool,                |")
	fmt.Println("| Welcome to our ICH,This is a good place to grow up.            |")
	fmt.Println("| please submit issue at : https://github.com/ssont/ichWebpass . |")
	fmt.Println("+----------------------------------------------------------------+")
	fmt.Println()
	flag.Parse()
	if *tcpPort <= 0 || *tcpPort >= 65536 {
		log.Fatalln("请输入正确的tcp端口。")
	}
	if *httpPort <= 0 || *httpPort >= 65536 {
		log.Fatalln("请输入正确的http端口。")
	}
	if *tcpPort == *httpPort {
		log.Fatalln("tcp端口与http端口不能为同一个。")
	}
	if *verifyKey == "" {
		log.Fatalln("必须输入一个验证的key")
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
retry:
	log.Println("客户端启动，连接：", *svrAddr, "， 端口：", *tcpPort, "， 并开启http服务端，端口为：", *httpPort)
	cli := NewRPClient(fmt.Sprintf("%s:%d", *svrAddr, *tcpPort), *httpPort)
	if err := cli.Start(); err != nil {
		log.Println(err)
		// 重连
		log.Println("5秒后重新连接...")
		time.Sleep(time.Second * 5)
		goto retry
	}
	defer cli.Close()
}
