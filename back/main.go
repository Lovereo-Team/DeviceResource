package main

import (
	adminRouters "DeviceResource/admin/routers"
	admin "DeviceResource/admin/service"
	"DeviceResource/config"
	"DeviceResource/core"
	"DeviceResource/core/response"
	genRouters "DeviceResource/generator/routers"
	gen "DeviceResource/generator/service"
	"DeviceResource/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

// initDI 初始化DI
func initDI() {
	regFunctions := admin.InitFunctions
	regFunctions = append(regFunctions, gen.InitFunctions...)
	regFunctions = append(regFunctions, core.GetDB)
	for i := 0; i < len(regFunctions); i++ {
		if err := core.ProvideForDI(regFunctions[i]); err != nil {
			log.Fatalln(err)
		}
	}
}

// initRouter 初始化router
func initRouter() *gin.Engine {
	// 初始化gin
	gin.SetMode(config.Config.GinMode)
	router := gin.New()
	// 设置静态路径
	router.Static(config.Config.PublicPrefix, config.Config.UploadDirectory)
	router.Static(config.Config.StaticPath, config.Config.StaticDirectory)
	// 设置中间件
	router.Use(gin.Logger(), middleware.Cors(), middleware.ErrorRecover())
	// 演示模式
	if config.Config.DisallowModify {
		router.Use(middleware.ShowMode())
	}
	// 特殊异常处理
	router.NoMethod(response.NoMethod)
	router.NoRoute(response.NoRoute)
	// 注册路由
	group := router.Group("/api")
	//core.RegisterGroup(group, routers.CommonGroup, middleware.TokenAuth())
	//core.RegisterGroup(group, routers.MonitorGroup, middleware.TokenAuth())
	//core.RegisterGroup(group, routers.SettingGroup, middleware.TokenAuth())
	//core.RegisterGroup(group, routers.SystemGroup, middleware.TokenAuth())

	routers := adminRouters.InitRouters[:]
	routers = append(routers, genRouters.InitRouters...)
	for i := 0; i < len(routers); i++ {
		core.RegisterGroup(group, routers[i])
	}
	return router
}

// initServer 初始化server
func initServer(router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           ":" + strconv.Itoa(config.Config.ServerPort),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func main() {
	// 刷新日志缓冲
	defer core.Logger.Sync()
	// 程序结束前关闭数据库连接
	if core.GetDB() != nil {
		db, _ := core.GetDB().DB()
		defer db.Close()
	}
	// 初始化DI
	initDI()
	// 初始化router
	router := initRouter()
	// 初始化server
	s := initServer(router)
	// 运行服务
	log.Fatalln(s.ListenAndServe().Error())
}

//package main
//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//)
//
//func main() {
//	// MySQL数据库连接信息
//	db, err := sql.Open("mysql", "root:120812614@tcp(localhost:3306)/DeviceResource?charset=utf8mb4&parseTime=True&loc=Local")
//	if err != nil {
//		panic(err.Error())
//	}
//	defer db.Close()
//
//	// 执行查询
//	rows, err := db.Query(`
//        SELECT table_name, TABLE_COMMENT, create_time, update_time
//        FROM information_schema.tables
//        WHERE table_schema = "DeviceResource"
//        AND table_name NOT LIKE "qrtz_%"
//        AND table_name NOT LIKE "gen_%"
//        AND table_name NOT IN (SELECT TABLE_NAME FROM la_gen_table) LIMIT 10`)
//	if err != nil {
//		panic(err.Error())
//	}
//	defer rows.Close()
//
//	// 遍历结果集并打印结果
//	for rows.Next() {
//		var tableName, tableComment, createTime string
//		var updateTime sql.NullString
//		err := rows.Scan(&tableName, &tableComment, &createTime, &updateTime)
//		if err != nil {
//			panic(err.Error())
//		}
//		fmt.Printf("Table Name: %s, Table Comment: %s, Create Time: %s", tableName, tableComment, createTime)
//		if updateTime.Valid {
//			fmt.Printf(", Update Time: %s\n", updateTime.String)
//		} else {
//			fmt.Printf(", Update Time: NULL\n")
//		}
//	}
//	if err := rows.Err(); err != nil {
//		panic(err.Error())
//	}
//}
