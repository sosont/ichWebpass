package main

import (
	"fmt"
	// "html"
	// "io"
	// "strings"
	"log"
	"net/http"
	"strconv"
	//	"time"

	"github.com/gin-gonic/gin"
)

func rateLimit(c *gin.Context) {
	ip := c.ClientIP()
	value := int(ips.Add(ip, 1))
	if value%50 == 0 {
		fmt.Printf("ip: %s, count: %d\n", ip, value)
	}
	if value >= 200 {
		if value%200 == 0 {
			fmt.Println("ip blocked")
		}
		c.Abort()
		c.String(503, "you were automatically banned :)")
	}
}

func index(c *gin.Context) {
	c.Redirect(301, "/login/index")
}

//登陆
func loginIndex(c *gin.Context) {
	c.HTML(200, "login.tpl.html", nil)
}

func checkLogin(c *gin.Context) {
	_u, _p := c.PostForm("username"), c.PostForm("password")
	log.Println(_u, "login:", _p)
	if _u == "admin" && _p == "ichadmin" {
		c.JSON(http.StatusOK, gin.H{"messages": "登陆成功", "success": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"messages": _u + "登陆失败:" + _p, "success": false})
	}
}

//主页
func IndexHome(c *gin.Context) {
	c.HTML(200, "index.tpl.html", nil)
}
func NatList(c *gin.Context) {
	c.HTML(200, "natlist.tpl.html", nil)
}

func startNatpass(c *gin.Context) {
	_t, _h := c.Query("tcpPort"), c.Query("httpPort")
	tp, err := strconv.Atoi(_t)
	hp, err := strconv.Atoi(_h)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"code": "ok", "tcp": tp, "http": hp})
	go startServer(tp, hp)
}

//获取数据记录
func getNatList() {

}
