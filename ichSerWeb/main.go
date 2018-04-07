package main

import (
	"database/sql"
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var vKey string = "Rm8FNu0xFY="

var verifyKey *string = &vKey

var db *sql.DB

func main() {
	ConfigRuntime()
	StartWorkers()
	StartWeb()
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func StartWorkers() {
	go statsWorker()
}

func StartWeb() {
	db, err := sql.Open("mysql", "root:yaoqi1717@tcp(139.199.177.131:13306)/yq_godb")
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(rateLimit, gin.Recovery()) //防重复提交频率
	router.LoadHTMLGlob("resources/*.tpl.html")
	router.Static("/static", "resources/static")
	router.GET("/", index)

	router.GET("/login/index", loginIndex)
	router.GET("/index/home", IndexHome)
	router.GET("/page/natlist", NatList)

	router.POST("/api/login", checkLogin)
	router.POST("/api/startNat", startNatpass)
	router.POST("/api/getNatlist/", getNatlist)

	router.Run(":9090")
}
