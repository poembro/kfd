package main

import ( 
    "log"
    "net/http" 
    //"reflect" 
    "kfd/internal/logic/db"
    "kfd/pkg/util"
    c "kfd/api/logic/http" 
)

func main() {
    defer util.RecoverPanic() 

    // 初始化数据库
    db.InitDB()

    // 初始化自增id配置
    util.InitUID(db.DBCli)
     
    http.HandleFunc("/api/login/wxlogin", c.AuthWxLogin)
    http.HandleFunc("/api/auth/login", c.AuthLogin)
    http.HandleFunc("/api/auth/register", c.AuthRegister)


    http.Handle("/api/user/info", c.Middleware(http.HandlerFunc(c.UserInfo)))
    http.Handle("/api/user/list", c.Middleware(http.HandlerFunc(c.UserList)))
    
    err := http.ListenAndServe("0.0.0.0:8080", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
        //panic(err)
    }
}