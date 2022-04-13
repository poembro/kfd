golang 商城 RESTful API接口服务
==============
一个快速入门 golang RESTful API 类型的商城项目 (精简版)
---------------------------------------

## 介绍
- 适用于刚入门golang的新手朋友，简单商城项目可以拿去直接用


## 功能点
- 完全基于 golang 官方的 net/http 实现接口 
- 对MySQL，redis 的增删改查
- net/http中间件 (用context保存当次请求数据，以便后续handler使用)
- token 验证授权 (openssl RSA 加解密)
- 列表接口,实现条件查询
- 详情页接口查询
- 精简的代码布局 控制器controller层 数据dao层 缓存cache层 服务service层
- 唯一id生成器，采用mysql表字段 控制


## 安装 (linux版本)
```
wget  https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
tar zxvf go1.13.4.linux-amd64.tar.gz 
mv go /usr/local/go

vi /etc/profile
export GOROOT=/usr/local/go
export GOPATH=/data/web/golang
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

go env -w GOPROXY=https://goproxy.cn,direct

git clone git@github.com:poembro/kfd.git
cd kfd
go build
```


##  感谢
- 代码参考了[gim](https://github.com/alberliu/gim), 感谢alberliu开源这么美的代码



--------------


# golang net/http 源码执行流程


## net/http 默认情况下的不足 
- 不能单独的对请求方法(POST,GET等)注册特定的处理函数
- 不支持Path变量获取参数 

## 使用案例
``` 
使用案例一:
func main{
    http.HandleFunc("/api/goods/info", c.GoodsInfo)
    http.Handle("/api/order/list", c.Middleware(http.HandlerFunc(c.OrderList)))
    http.ListenAndServe("0.0.0.0:8080", nil) //设置监听的端口
}

使用案例二:
func main{
    type router struct {} 
    func (ro *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        switch r.URL.Path {
        case "/ping": 
            w.Write([]byte(`pong`))
        default:
            w.WriteHeader(http.StatusNotFound)
        }
    }
    err := http.ListenAndServe("0.0.0.0:8080", &router{}) //设置路由和监听的端口
}

使用案例三:
func main{
    http.HandleFunc("/api/test", c.GoodsInfo)
    http.Handle("/ping", c.Middleware(http.HandlerFunc(c.Ping)))
    srv := http.Server{
        Addr:    "0.0.0.0:8080",
        Handler: http.DefaultServeMux, //上面所有api路由(/ping,/api/test)全都在该全局结构体指针的m属性上
    }
    err := srv.ListenAndServe() //设置监听的端口
    ....
}
```

## 源码分析  

### 第一步: 注册路由
``` 
// 定义1个接口   (路由结构体必须实现 ServeHTTP 方法)
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

 // 这一步极为重要 开始就定义了1个全局"路由结构体指针", 这才有了 使用案例一 使用案例三 的操作
var DefaultServeMux = &defaultServeMux 
type ServeMux struct {
    mu    sync.RWMutex
    m     map[string]muxEntry
    es    []muxEntry // 从最长到最短排序的条目的切片。
    hosts bool // whether any patterns contain hostnames
}

var defaultServeMux ServeMux  //默认的 DefaultServeMux 路由的

func Handle(pattern string, handler Handler) { 
    DefaultServeMux.Handle(pattern, handler) 
}

func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    DefaultServeMux.HandleFunc(pattern, handler)  //DefaultServeMux.HandleFunc 函数最终会调用 ServeMux.Handle函数。
}

func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    mux.Handle(pattern, HandlerFunc(handler)) //为给定的模式注册处理程序函数
}

func (mux *ServeMux) Handle(pattern string, handler Handler) {
    //省略加锁和判断代码 
    //把我们注册的路径和相应的处理函数存入了m字段中
    mux.m[pattern] = muxEntry{h: handler, pattern: pattern} 
    if pattern[0] != '/' {
        mux.hosts = true
    }
}

func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
    //省略一些无关代码
    h, _ := mux.Handler(r) // h是1个新的结构体并且实现了接口 Handler
    h.ServeHTTP(w, r)
}
``` 


### 第二步: 创建 绑定 监听 
``` 
server.go
type Server struct {
    Addr string
    Handler Handler   // 这一步极为重要 路由结构体 被赋值过来
    .....
}

// 监听端口
//  func ListenAndServe(addr string, handler Handler) error {...} 等同于下面
func (srv *Server) ListenAndServe() error {
    ln, err := net.Listen("tcp", srv.Addr)
    return srv.Serve(ln)
}
// 接收连接请求 并读取连接中的头
func (srv *Server) Serve(l net.Listener) error {
    for {
        ....
        rw, err := l.Accept()
        c := srv.newConn(rw)
        go c.serve(connCtx) //执行 结构体 下的serve 方法
    }
} 
func (srv *Server) newConn(rwc net.Conn) *conn {
    c := &conn{ 
        server: srv, // 创建1个连接结构体,  srv赋值进去了.   注意srv.Handler.m这个map里面放的是注册好的路由
        rwc:    rwc,
    }
    return c
}

// 处理一个新的连接
conn 结构体 
func (c *conn) serve(ctx context.Context) {
    for {
        w, err := c.readRequest(ctx) //return *response, error
        serverHandler{c.server}.ServeHTTP(w, w.req) // 这一步极为重要 表示程序源码结尾
    }
}

``` 


### 第三步：源码结尾 (路由结构体的 ServeHTTP 方法被调用) 
``` 
type serverHandler struct {
    srv *Server  
}
func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
    handler := sh.srv.Handler // 注意 注意 注意 handler 这个就路由结构体, srv 对应第二步 Server 结构体
    if handler == nil {
        handler = DefaultServeMux  // 这个变量是不是很熟悉 就是 第一步 中的全局路由结构体指针
    }
    if req.RequestURI == "*" && req.Method == "OPTIONS" {
        handler = globalOptionsHandler{}
    }
    handler.ServeHTTP(rw, req) //这里则是 调用了第一步 路由结构体的 ServeHTTP 方法
}

```
