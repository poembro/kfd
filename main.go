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

    http.HandleFunc("/api/goods/list", c.GoodsList)
    http.HandleFunc("/api/goods/info", c.GoodsInfo)

    http.HandleFunc("/api/goodscategory/list", c.GoodsCategoryList)
    http.HandleFunc("/api/goodscategory/info", c.GoodsCategoryInfo)

    http.Handle("/api/order/list", c.Middleware(http.HandlerFunc(c.OrderList)))
    http.Handle("/api/order/info", c.Middleware(http.HandlerFunc(c.OrderInfo)))
    http.Handle("/api/order/save", c.Middleware(http.HandlerFunc(c.OrderSave)))

    http.Handle("/api/user/info", c.Middleware(http.HandlerFunc(c.UserInfo)))
    http.Handle("/api/user/list", c.Middleware(http.HandlerFunc(c.UserList)))

    err := http.ListenAndServe("0.0.0.0:8080", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
        //panic(err)
    }
}