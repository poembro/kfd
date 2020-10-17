golang web 入门
==============
一个快速实战 golang restful api 类型的项目

---------------------------------------

## 介绍
- 完全基于 golang 原始 net/http 实现接口 
- 对MySQL，redis 的增删改查
- 基于[gim](https://github.com/alberliu/gim)编写, 感谢[gim](https://github.com/alberliu/gim)作者开源这么美的代码


## 已经实现 API
- 中间件 (用context保存当次请求数据，以便后续handler使用)
- 验证授权 (openssl RSA 加解密)
- 列表条件查询
- 详情页查询
- 层次分明 控制器controller层 数据dao层 缓存cache层 服务service层
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


