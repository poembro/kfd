
### 注册接口
[POST] http://127.0.0.1:8080/auth/register
- 注意content-type : application/json

| Name            | Type     | test                 |
|:----------------|:--------:|:-----------------------|
| [account]:账号  | string    | 13200000000 |
| [password]:密码 | string    | 13200000000 |
 
response:
```
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
          "account": "13200000000",
          "Password": "aeb3f03a7828263b54ae16c7149947d7"
     }
}
```

### 登陆接口
[POST] http://127.0.0.1:8080/auth/login
- 注意content-type : application/json

| Name            | Type     | test                 |
|:----------------|:--------:|:-----------------------|
| [account]:账号  | string    | 13200000000 |
| [password]:密码 | string    | 13200000000 |
 
response:
```
{
  "errCode": 0,
  "msg": "success",
  "data": "xLTDb/Ph7zkdFfzNFxlW+R47tW34FDJqzeHbEhPsULDv8hWUew+..."
}
```



### 列表接口
[POST] http://127.0.0.1:8080/user/list
- header['token'] = "登陆接口返回的data值"
- 注意content-type : application/json 

| Name            | Type     | test                 |
|:----------------|:--------:|:-----------------------|
| [sex]:男/女  | int32    | 2 |
| [currentpage]:当前页 | int    | 1 |
| [pagesize]:页长 | int    | 2 |

response:
```
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
[POST] http://127.0.0.1:8080/user/info
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