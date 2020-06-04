package config

import (
	"os"
	"net/url"
)

func init() {
	env := os.Getenv("kfd_env")
	switch env {
	case "prod":
		initProdConf()
	default:
		initLocalConf()
	}
}

// logic配置
type logicConf struct {
	MySQL                  string 
	RedisIP                string
	ApiHTTP                string
 	ApiRPC                 string
	ConnRPCAddrs           string
}

type weixinConf struct {
	SappJscodeURL          string 
	Appid                  string
	Secret                 string 
}



var (
	LogicConf logicConf
	WinxinConf weixinConf
)

func initLocalConf() {
	loc := url.QueryEscape("Asia/Shanghai")
	LogicConf = logicConf{
		MySQL: "root:admin@2019@tcp(sh-cdb-dw5p0dpo.sql.tencentcdb.com:61740)/db_scmj?charset=utf8&loc="+loc+"&parseTime=true",
		RedisIP: "127.0.0.1:6379",
		ApiHTTP: "0.0.0.0:8080", 
		ApiRPC: "0.0.0.0:50000", 
		ConnRPCAddrs: "addrs:///127.0.0.1:50000,127.0.0.1:50001",
	}

	WinxinConf = weixinConf{
		SappJscodeURL: "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		Appid: "wx26d393bcc04ccccfb",
        Secret: "fa7c9a2044a768d28c0caa3a7d0a70813a",
	}
}

func initProdConf() {
	LogicConf = logicConf{
		MySQL: "root:liu123456@tcp(localhost:3306)/gim?charset=utf8&parseTime=true",
		RedisIP: "127.0.0.1:6379",
		ApiHTTP: "0.0.0.0:8080", 
		ApiRPC: "0.0.0.0:50000", 
		ConnRPCAddrs: "addrs:///127.0.0.1:50000,127.0.0.1:50001",
	}
}