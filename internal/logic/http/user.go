package http

import (
    "context"
    "net/http"
    "io/ioutil"  
    json "encoding/json" 
    "kfd/internal/logic/model" 
    service "kfd/internal/logic/service"
)

func UserInfo(w http.ResponseWriter, r *http.Request) {
    var (
        app_id int64 = 1
        user_id int64
        ok bool
    )
    //user_id = r.Header.Get("User-Agent")
    user_id, ok = r.Context().Value("user_id").(int64); //作为赋值中的类型字符串:需要类型断言
    if user_id == 0 || !ok {
        OutJson(w, 1, "token error1 r.Context().Value ", user_id)
        return  
    }

    user, err := service.UserService.Get(context.TODO(), app_id, user_id)
    if err != nil  {
        OutJson(w, 1, "error user not register", user_id)
        return
    }
    
    OutJson(w, 0, "success", user)
    return
}

type Page struct{
    Sex            int32  `json:"sex"`           // sex
    Count          int    `json:"count"`         // count
    TotalPages     int    `json:"totalpages"`    // totalpages
    PageSize       int    `json:"pagesize"`      // pagesize
    CurrentPage    int    `json:"currentpage"`   // currentpage
    Data           []model.User `json:"data"`
}

func UserList(w http.ResponseWriter, r *http.Request) {
    var (
        param Page
    )
    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &param)
    
    if  param.PageSize <= 0  {
        param.PageSize = 10
    }
    if  param.CurrentPage <= 0  {
        param.CurrentPage = 1
    }

    count, user, err := service.UserService.ListUser(context.TODO(), 
        param.Sex, param.CurrentPage, param.PageSize)
    if err != nil  {
        OutJson(w, 1, "error user ", err)
        return
    }
    param.Count = count
    param.TotalPages = TotalPage(count, param.PageSize)
    param.Data = user

    OutJson(w, 0, "success", param)
    return
}


