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