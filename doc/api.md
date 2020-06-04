
### 注册接口
[POST] http://127.0.0.1:8080/api/auth/register
- 注意content-type : application/json

| Name            | Type     | test                 |
|:----------------|:--------:|:-----------------------|
| [account]:账号  | string    | 13200000000 |
| [password]:密码 | string    | 13200000000 |

```
request 
{"account":"13200000002", "password":"13200000002"}

response: 
{
     "errCode": 0,
     "msg": "success",
     "data": {
          "app_id": 1,
          "user_id": 1831,
          "nickname": "昵称",
          "sex": 1,
          "avatar_url": "/#",
          "extra": "#",
          "account": "13200000002",
          "Password": "aeb3f03a7828263b54ae16c7149947d7"
     }
}
```
 
### 登陆接口
[POST] http://127.0.0.1:8080/api/auth/login
- 注意content-type : application/json

| Name            | Type     | test                 |
|:----------------|:--------:|:-----------------------|
| [account]:账号  | string    | 13200000000 |
| [password]:密码 | string    | 13200000000 |
 
```
request 
{"account":"13200000002", "password":"13200000002"}

response:
{
  "errCode": 0,
  "msg": "success",
  "data": "xLTDb/Ph7zkdFfzNFxlW+R47tW34FDJqzeHbEhPsULDv8hWUew+..."
}
```

 

### 列表接口
[POST] http://127.0.0.1:8080/api/user/list
- header['token'] = "登陆接口返回的data值"
- 注意content-type : application/json 

| Name            | Type     | test                 |
|:----------------|:--------:|:-----------------------|
| [sex]:男/女  | int32    | 2 |
| [currentpage]:当前页 | int    | 1 |
| [pagesize]:页长 | int    | 2 |

```
request 
{"currentpage":1, "pagesize":10, "sex":2}

response:
{
     "errCode": 0,
     "msg": "success",
     "data": {
          "sex": 2,
          "count": 0,
          "totalpages": 0,
          "pagesize": 2,
          "currentpage": 1,
          "data": [
               {
                    "app_id": 0,
                    "user_id": 3,
                    "nickname": "3",
                    "sex": 2,
                    "avatar_url": "avatar_url",
                    "extra": "extra",
                    "account": "",
                    "Password": ""
               },
               {
                    "app_id": 0,
                    "user_id": 2,
                    "nickname": "2",
                    "sex": 2,
                    "avatar_url": "avatar_url",
                    "extra": "extra",
                    "account": "",
                    "Password": ""
               }
          ]
     }
}
```



### 详细接口
[POST] http://127.0.0.1:8080/api/user/info
- header['token'] = "登陆接口返回的data值"
- 注意content-type : application/json 

| Name            | Type     | test                 |
|:----------------|:--------:|:-----------------------|
| --  | --    | -- | 

response:
```
 {
     "errCode": 0,
     "msg": "success",
     "data": {
          "app_id": 0,
          "user_id": 1711,
          "nickname": "昵称",
          "sex": 1,
          "avatar_url": "/#",
          "extra": "#",
          "account": "13200000000",
          "Password": "aeb3f03a7828263b54ae16c7149947d7"
     }
}
```




### 商品列表接口
[POST] http://127.0.0.1:8080/api/goods/list
- 注意content-type : application/json 

| Name            | Type     | test                 |
|:----------------|:--------:|:-----------------------|
| [category_id]:商品分类  | int32    | 2 |
| [currentpage]:当前页 | int    | 1 |
| [pagesize]:页长 | int    | 2 |

```
request 
{"currentpage":1, "pagesize":10, "category_id":1}

response:
{
     "errcode": 0,
     "msg": "success",
     "data": {
          "category_id": 1,
          "count": 2,
          "totalpages": 1,
          "pagesize": 10,
          "currentpage": 1,
          "data": [
               {
                    "id": 2,
                    "name": "三峡大坝一日游",
                    "description": "1天1晚不一样的“疫”样的时光",
                    "category_id": 1,
                    "is_on_sale": 1,
                    "thumb": "https://file-1259574925.cos.ap-shanghai.myqcloud.com/nste-img/201911/5523231269812338_640-1.jpeg",
                    "goods_price": 800,
                    "market_price": 1000,
                    "sells_num": 0,
                    "comments_num": 1,
                    "favorites_num": 1,
                    "is_hot": 1,
                    "goods_sort": 1,
                    "inventory": 1,
                    "visible": 11,
                    "create_time": "",
                    "update_time": ""
               },
               {
                    "id": 1,
                    "name": "七天六晚不一样的“疫”样的时光",
                    "description": "七天六晚不一样的“疫”样的时光",
                    "category_id": 1,
                    "is_on_sale": 1,
                    "thumb": "https://file-1259574925.cos.ap-shanghai.myqcloud.com/nste-img/201911/5523231269812338_640-1.jpeg",
                    "goods_price": 1800,
                    "market_price": 2000,
                    "sells_num": 0,
                    "comments_num": 1,
                    "favorites_num": 1,
                    "is_hot": 1,
                    "goods_sort": 1,
                    "inventory": 1,
                    "visible": 11,
                    "create_time": "",
                    "update_time": ""
               }
          ]
     }
}
```




### 商品详细接口
[POST] http://127.0.0.1:8080/api/goods/info
- 注意content-type : application/json 

| Name            | Type     | test                 |
|:----------------|:--------:|:-----------------------|
| [id]:商品id  | int32    | 2 |
```
request 
{"id":1}

response:
{
     "errcode": 0,
     "msg": "success",
     "data": {
          "id": 0,
          "name": "七天六晚不一样的“疫”样的时光",
          "description": "七天六晚不一样的“疫”样的时光",
          "category_id": 1,
          "is_on_sale": 1,
          "thumb": "https://file-1259574925.cos.ap-shanghai.myqcloud.com/nste-img/201911/5523231269812338_640-1.jpeg",
          "goods_price": 1800,
          "market_price": 2000,
          "sells_num": 0,
          "comments_num": 1,
          "favorites_num": 1,
          "is_hot": 1,
          "goods_sort": 1,
          "inventory": 1,
          "visible": 11,
          "create_time": "",
          "update_time": ""
     }
}
```





### 下单接口
[POST] http://127.0.0.1:8080/api/order/save
- header['token'] = "登陆接口返回的data值"
- 注意content-type : application/json 

| Name            | Type     | test                 |
|:----------------|:--------:|:-----------------------|
| [goods_id]:商品id  | int32    | 2 |
| [goods_num]:商品数量  | int32    | 2 |
```
request 
[{"goods_id":1, "goods_num":1},{"goods_id":2, "goods_num":1}]

response:
{
     "errcode": 0,
     "msg": "success",
     "data": {
          "id": 0,
          "name": "七天六晚不一样的“疫”样的时光",
          "description": "七天六晚不一样的“疫”样的时光",
          "category_id": 1,
          "is_on_sale": 1,
          "thumb": "https://file-1259574925.cos.ap-shanghai.myqcloud.com/nste-img/201911/5523231269812338_640-1.jpeg",
          "goods_price": 1800,
          "market_price": 2000,
          "sells_num": 0,
          "comments_num": 1,
          "favorites_num": 1,
          "is_hot": 1,
          "goods_sort": 1,
          "inventory": 1,
          "visible": 11,
          "create_time": "",
          "update_time": ""
     }
}
```



### 订单列表接口
[POST] http://127.0.0.1:8080/api/goods/list
- 注意content-type : application/json 

| Name            | Type     | test                 |
|:----------------|:--------:|:-----------------------|
| [category_id]:商品分类  | int32    | 2 |
| [currentpage]:当前页 | int    | 1 |
| [pagesize]:页长 | int    | 2 |

```
request 
{"currentpage":1, "pagesize":10, "status":1}

response:
{
     "errcode": 0,
     "msg": "success",
     "data": {
          "uid": 0,
          "status": 1,
          "count": 53,
          "totalpages": 53,
          "pagesize": 1,
          "currentpage": 1,
          "data": [
               {
                    "Order": {
                         "id": 0,
                         "order_sn": "O_202006041209",
                         "uid": 2802,
                         "status": 1,
                         "payid": 1,
                         "paytime": 0,
                         "ispay": 1,
                         "shiptime": 0,
                         "finished_time": 0,
                         "refund_msg": "",
                         "cancel_type": 0,
                         "cancel_remark": "",
                         "realname": "昵称",
                         "mobile": "",
                         "shopname": "",
                         "address": "#",
                         "total_price": 2600,
                         "remark": "",
                         "visible": 1,
                         "create_time": "2020-06-04 16:47:52",
                         "update_time": "2020-06-04 16:47:52"
                    },
                    "OrderGoods": [
                         {
                              "id": 119,
                              "order_sn": "O_202006041209",
                              "goods_id": 2,
                              "refund_msg": "",
                              "cancel_type": 0,
                              "cancel_remark": "",
                              "goods_name": "三峡大坝一日游",
                              "thumb": "https://file-1259574925.cos.ap-shanghai.myqcloud.com/nste-img/201911/5523231269812338_640-1.jpeg",
                              "goods_num": 1,
                              "goods_price": 800,
                              "goods_spec": 0,
                              "status": 0,
                              "visible": 1,
                              "create_time": "2020-06-04 16:47:52",
                              "update_time": "2020-06-04 16:47:52"
                         },
                         {
                              "id": 118,
                              "order_sn": "O_202006041209",
                              "goods_id": 1,
                              "refund_msg": "",
                              "cancel_type": 0,
                              "cancel_remark": "",
                              "goods_name": "七天六晚不一样的“疫”样的时光",
                              "thumb": "https://file-1259574925.cos.ap-shanghai.myqcloud.com/nste-img/201911/5523231269812338_640-1.jpeg",
                              "goods_num": 1,
                              "goods_price": 1800,
                              "goods_spec": 0,
                              "status": 0,
                              "visible": 1,
                              "create_time": "2020-06-04 16:47:52",
                              "update_time": "2020-06-04 16:47:52"
                         }
                    ]
               }
          ]
     }
}
```



### 订单详细接口
[POST] http://127.0.0.1:8080/api/goods/info
- 注意content-type : application/json 

| Name            | Type     | test                 |
|:----------------|:--------:|:-----------------------|
| [order_sn]:商品id  | int32    | 2 |
```
request 
{"order_sn":"O_202006041209"}

response:
{
     "errcode": 0,
     "msg": "success",
     "data": {
          "id": 0,
          "name": "七天六晚不一样的“疫”样的时光",
          "description": "七天六晚不一样的“疫”样的时光",
          "category_id": 1,
          "is_on_sale": 1,
          "thumb": "https://file-1259574925.cos.ap-shanghai.myqcloud.com/nste-img/201911/5523231269812338_640-1.jpeg",
          "goods_price": 1800,
          "market_price": 2000,
          "sells_num": 0,
          "comments_num": 1,
          "favorites_num": 1,
          "is_hot": 1,
          "goods_sort": 1,
          "inventory": 1,
          "visible": 11,
          "create_time": "",
          "update_time": ""
     }
}
```