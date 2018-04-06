package main

import (
	"flag"
	"log"
)

var (
	tcpPort   = flag.Int("tp", 0, "Socket连接或者监听的端口")
	httpPort  = flag.Int("hp", 0, "当mode为server时为服务端监听端口，当为mode为client时为转发至本地客户端的端口")
	svrAddr   = flag.String("addr", "127.0.0.1", "为连接服务器的地址,考虑取消")
	verifyKey = flag.String("vkey", "", "用作客户端与服务端连接时的校验")
)

func main_test() {
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
	log.Println("服务端启动，监听tcp服务端端口：", *tcpPort, "， http服务端端口：", *httpPort)
	svr := NewRPServer(*tcpPort, *httpPort)
	if err := svr.Start(); err != nil {
		log.Fatalln(err)

		defer svr.Close()
	}
}
