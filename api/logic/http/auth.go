package http

import (
    "fmt"
    "context"
    "net/http"
    "io/ioutil" 
    "encoding/json"
    "kfd/config"
    "kfd/pkg/util" 
    "kfd/internal/logic/model" 
    service "kfd/internal/logic/service" 
)

 
func toRemoteGetOpenid(js_code string) (arr []byte) { 
    var ( 
		url string
	) 
    url = fmt.Sprintf(config.WinxinConf.SappJscodeURL, config.WinxinConf.Appid, config.WinxinConf.Secret, js_code)
	arr = util.HttpGet(url) 
	return
}

type AuthLoginBody struct {
    Code     string `json:"code"`
    Nickname string `json:"nickname"` 
    Face     string `json:"face"` 
}

type Wxinfo struct {
	Session_key string `json:"session_key"`
    Openid      string `json:"openid"`
} 
//微信登录接口
func AuthWxLogin(w http.ResponseWriter, r *http.Request) {
    var param AuthLoginBody
    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &param)

    var wx Wxinfo
    tmp := toRemoteGetOpenid(param.Code) 
    json.Unmarshal(tmp, &wx)
    if wx.Openid == "" {
        OutJson(w, 1, " Openid error", wx)
        return
    }
    //通过获取到的openid 判断是否已存在
    already, err := service.UserService.GetUserthird(context.TODO(), 1, wx.Openid)
    if already != nil {
        token, err := service.AuthService.BuildToken(context.TODO(), already.UserId, "13200000000")
        if err != nil {
            OutJson(w, 1, "token error", token)
            return
        }
        OutJson(w, 0, "success", token) 
        return
    }
    
    //id生成器获取id
    UserId, err := util.DeviceIdUid.Get()
    if err != nil {
        OutJson(w, 1, "id error", nil)
        return
    }
    //写入用户表
    addUser := model.User {
        AppId : 1,
        UserId : UserId,
        Nickname : param.Nickname,
        Sex : 1,
        AvatarUrl : param.Face,
        Extra : "#", 
        Account : "",
        Password : "",
    }
    err = service.UserService.Add(context.TODO(), addUser)
    if err != nil {
        OutJson(w, 1, "sql error ", err)
        return
    }
    //写入第三方登录表
    addUserthirdArr := model.Userthird {  UserId: UserId, Typeid: 1, Openid: wx.Openid, }
    service.UserService.AddUserthird(context.TODO(), addUserthirdArr)

    //构造token
    token, err := service.AuthService.BuildToken(context.TODO(), UserId, "13200000000")
    if err != nil {
        OutJson(w, 1, "error", token)
        return
    }
    OutJson(w, 0, "success", token)
    return
}

//普通登录接口
func AuthLogin(w http.ResponseWriter, r *http.Request) {
    var (
        err error
        token string
        param model.User
    )
    
    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &param)
 
    if param.Account == "" {
        OutJson(w, 1, "error not Account", nil)
        return
    }
    if param.Password == "" {
        OutJson(w, 1, "error not Password", nil)
        return
    }
    
    user, _ := service.UserService.AccountByUser(context.TODO(), param.Account)
    if user == nil  {
        OutJson(w, 1, "error user not register", err)
        return
    }

    password := util.Md5(param.Password + "123456")
    if user.Password != password {
        OutJson(w, 1, "password error " + password, user)
        return
    }
    
    token, err = service.AuthService.BuildToken(context.TODO(), user.UserId, user.Account)
    if err != nil {
        OutJson(w, 1, "error", nil)
        return
    }
    OutJson(w, 0, "success", token)
    return
}

//普通注册接口
func AuthRegister(w http.ResponseWriter, r *http.Request) {
    var (
        param model.User 
    )
    
    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &param)
 
    if param.Account == "" {
        OutJson(w, 1, "error not Account", nil)
        return
    }
    if param.Password == "" {
        OutJson(w, 1, "error not Password", nil)
        return
    }
    
    UserId, err := util.DeviceIdUid.Get()
    if err != nil {
        OutJson(w, 1, "error", nil)
        return
    }
    
    param.UserId = UserId
    param.Password = util.Md5(param.Password + "123456")
    param.AppId = 1
    param.Nickname = "昵称"
    param.Sex = 1
    param.AvatarUrl = "#"
    param.Extra = "#"
  
    err = service.UserService.Add(context.TODO(), param)
    if err != nil {
        OutJson(w, 1, "sql error ", err)
        return
    }

    param.Password = ""
    OutJson(w, 0, "success", param)
}
