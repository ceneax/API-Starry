package config

import (
	"Starry/util"
	"os"
	"path/filepath"
	"time"

	"github.com/go-ini/ini"
)

/*
 * 全局变量
 * 读取配置文件获取值
 */
var (
	Cfg *ini.File // config配置文件

	RunMode string // 运行模式：debug | release

	ServerHost     string // 服务端运行监听地址，包含端口
	ServerHttpPort int    // http服务的端口

	LogReadTimeout  time.Duration // 读取超时时间
	LogWriteTimeout time.Duration // 写入超时时间
	LogFilePath     string        // 日志文件路径
	LogFileName     string        // 日志文件名

	PageSize int // 服务端返回分页结果的每页条数

	IdentityKey string // 用户表的字段key

	DBType     string // 数据库类型
	DBUser     string // 数据库登录用户名
	DBPassword string // 数据库登录密码
	DBHost     string // 数据库地址
	DBPort     string // 数据库端口
	DBName     string // 数据库名
)

/*
 * 初始化
 */
func init() {
	// 获取工作路径
	workPath, err := os.Getwd()
	if err != nil {
		// 工作路径获取失败
		panic(err)
	}

	const fileName = "config.ini"
	// 生成config.ini配置文件的完整路径
	var filePath = filepath.Join(workPath, fileName)

	// 判断文件是否存在
	if !util.FileExists(filePath) {
		appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			panic(err)
		}

		filePath = filepath.Join(appPath, fileName)
		if !util.FileExists(filePath) {
			panic("初始化错误：配置文件config.ini不存在！")
		}
	}

	Cfg, err = ini.Load(filePath)
	if err != nil {
		panic("初始化错误：配置文件config.ini读取失败！原因：" + err.Error())
	}

	loadBase()
	loadServer()
	loadApp()
	loadDatabase()
}

/*
 * 加载基本参数
 */
func loadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

/*
 * 加载server节点参数
 */
func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		panic("初始化错误：配置文件的server节点数据读取失败！原因：" + err.Error())
	}

	ServerHost = sec.Key("SERVER_HOST").MustString("localhost:4444")
	ServerHttpPort = sec.Key("HTTP_PORT").MustInt(4445)
	LogReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	LogWriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
	LogFilePath = sec.Key("LOG_FILEPATH").MustString("log.txt")
}

/*
 * 加载app节点参数
 */
func loadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		panic("初始化错误：配置文件的app节点数据读取失败！原因：" + err.Error())
	}

	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	IdentityKey = sec.Key("IDENTITY_KEY").MustString("!@)*#)!@U#@*!@!)")
}

/*
 * 加载database节点参数
 */
func loadDatabase() {
	sec, err := Cfg.GetSection("database")
	if err != nil {
		panic("初始化错误：配置文件的database节点数据读取失败！原因：" + err.Error())
	}

	DBType = sec.Key("TYPE").MustString("mysql")
	DBUser = sec.Key("USER").MustString("test")
	DBPassword = sec.Key("PASSWORD").MustString("test")
	DBHost = sec.Key("HOST").MustString("127.0.0.1")
	DBPort = sec.Key("PORT").MustString("3306")
	DBName = sec.Key("NAME").MustString("test")
}
