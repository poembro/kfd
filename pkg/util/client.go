package util

import (
    "net/http"
    "io"
    "bytes"
    "encoding/json"
    "io/ioutil"
    "time"
)
var client *http.Client
func init() {
    client = &http.Client{
        Transport: &http.Transport{
            MaxIdleConns:100,
            MaxIdleConnsPerHost:2,
        },
        Timeout: time.Duration(5 * time.Second),
    }
}

//发送GET请求
//url:请求地址
//res:请求返回的内容
func HttpGet(url string) (res []byte) {
    resp, error := client.Get(url)
    defer resp.Body.Close()
    if error != nil {
        panic(error)
    }

    var buffer [512]byte
    buf := bytes.NewBuffer(nil)
    for {
        n, err := resp.Body.Read(buffer[0:])
        buf.Write(buffer[0:n])
        if err != nil && err == io.EOF {
            break
        } else if err != nil {
            panic(err)
        }
    }

    res = buf.Bytes()
    return
}

//发送POST请求
//url:请求地址，data:POST请求提交的数据,contentType:请求体格式，如：application/json
//content:请求放回的内容
func HttpPost(url string, data interface{}) (content string) {
    jsonStr, _ := json.Marshal(data)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr)) //也可以 http.POST(url,body) 因为POST方法里面也是 http.NewRequest 毫无意义的包装了一下而已
    if err != nil {
        panic(err)
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("Cookie", "name=anny") 
    defer req.Body.Close()
 
    resp, error := client.Do(req)
    if error != nil {
        panic(error)
    }
    defer resp.Body.Close()

    result, _ := ioutil.ReadAll(resp.Body)
    content = string(result)
    return
}

//看了下http源码
/*
//得到request接构体   
//随便复习了一下context是个接口里面有 Done() <-chan struct{}
// Err() error
//Deadline() (deadline time.Time, ok bool)
//Value(key interface{}) interface()

## 第一步 必须要有个 Request 结构变量 像http.POST http.GET 方法 本质上就是包装此函数
func NewRequest(method, url string, body io.Reader) (*Request, error) {
	return NewRequestWithContext(context.Background(), method, url, body)
}
func NewRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*Request, error) 
type Request struct { 
	Method string 
	URL *url.URL 
	Proto      string // "HTTP/1.0"
	ProtoMajor int    // 1
	ProtoMinor int    // 0 
	Header Header 
	Body io.ReadCloser     注意这里接收的是 io.Reader
	GetBody func() (io.ReadCloser, error)   注意这里都会包一个函数里面不管body是啥(*bytes.Buffer/*bytes.Reader/*strings.Reader)类型都转换   io.ReadCloser
	ContentLength int64 
	TransferEncoding []string 
	Close bool 
	Host string 
	Form url.Values 
	PostForm url.Values 
	MultipartForm *multipart.Form 
	Trailer Header 
	RemoteAddr string 
	RequestURI string 
	TLS *tls.ConnectionState 
	Cancel <-chan struct{} 
	Response *Response 
	ctx context.Context
}

## 第二步 发送连接网络发送请求数据
特别注意 client.go中没有发送步骤  在transport.go文件中 
https://studygolang.com/articles/5774



net/http 的不足 
不能单独的对请求方法(POST,GET等)注册特定的处理函数
不支持Path变量参数 
……

例如
func main{
    http.HandleFunc("/api/goods/info", c.GoodsInfo)
    http.Handle("/api/order/list", c.Middleware(http.HandlerFunc(c.OrderList)))
    http.ListenAndServe("0.0.0.0:8080", nil) //设置监听的端口
}

## 第一步: 注册路由
``` 
type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	hosts bool // whether any patterns contain hostnames
}

var defaultServeMux ServeMux  //默认的 DefaultServeMux 路由的
var DefaultServeMux = &defaultServeMux 

func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)  //DefaultServeMux.HandleFunc 函数最终会调用 ServeMux.Handle函数。
}

func (mux *ServeMux) Handle(pattern string, handler Handler) {
	//省略加锁和判断代码 
	if mux.m == nil {
		mux.m = make(map[string]muxEntry)
	}
	//把我们注册的路径和相应的处理函数存入了m字段中
	mux.m[pattern] = muxEntry{h: handler, pattern: pattern}

	if pattern[0] != '/' {
		mux.hosts = true
	}
}

func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	//省略一些无关代码
	h, _ := mux.Handler(r)
	h.ServeHTTP(w, r)
}
``` 


## 第二步: 创建 绑定 监听  while(1) accept() 
``` 
server.go
func ListenAndServe(addr string, handler Handler) error {
server.ListenAndServe()
ln, err := net.Listen("tcp", addr)
func (srv *Server) Serve(l net.Listener) error {

func (srv *Server) newConn(rwc net.Conn) *conn {
	c := &conn{
		server: srv,
		rwc:    rwc,
	}
	if debugServerConnections {
		c.rwc = newLoggingConn("server", c.rwc)
	}
	return c
}
go c.serve(connCtx) 

func (c *conn) serve(ctx context.Context) 

``` 


## 第三步：处理HTTP请求    会调用Handler接口的ServeHTTP方法，而ServeMux正好实现了Handler。
``` 
type serverHandler struct {
	srv *Server
}
func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
	handler := sh.srv.Handler
	if handler == nil {
		handler = DefaultServeMux
	}
	if req.RequestURI == "*" && req.Method == "OPTIONS" {
		handler = globalOptionsHandler{}
	}
	handler.ServeHTTP(rw, req)
}

serverHandler{conn.server}.ServeHTTP(w, w.req)
``` 

*/