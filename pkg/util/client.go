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
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
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