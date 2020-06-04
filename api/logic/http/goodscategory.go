package http

import (
    "context"
    "net/http"
    "io/ioutil"  
    json "encoding/json" 
    "kfd/internal/logic/model" 
    service "kfd/internal/logic/service"
)

func GoodsCategoryInfo(w http.ResponseWriter, r *http.Request) {
    var param = model.GoodsCategory{}
    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &param)
    
    if param.Id <= 0 {
        OutJson(w, 1, "goods id error ", param.Id)
        return 
    }
    
    item, err := service.GoodsCategoryService.Get(context.TODO(), param)
    if err != nil  {
        OutJson(w, 1, "error item not ", param.Id)
        return
    }
    
    OutJson(w, 0, "success", item)
    return
}

type GoodsCategoryPage struct{
    Name           string  `json:"name"`   // sex
    Count          int    `json:"count"`        // count
    TotalPages     int    `json:"totalpages"`   // count
    PageSize       int    `json:"pagesize"`     // count
    CurrentPage    int    `json:"currentpage"`  // count
    Data           []model.GoodsCategory
}


func GoodsCategoryList(w http.ResponseWriter, r *http.Request) {
    var (
        param GoodsCategoryPage
        opt model.GoodsCategory
    )
    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &param)
    
    if  param.PageSize <= 0  {
        param.PageSize = 10
    }
    if  param.CurrentPage <= 0  {
        param.CurrentPage = 1
    }

    opt.Name = param.Name
    count, arr, err := service.GoodsCategoryService.List(context.TODO(),  opt, param.CurrentPage, param.PageSize)
    if err != nil  {
        OutJson(w, 1, "error arr ", err)
        return
    }
    param.Count = count
    param.TotalPages = TotalPage(count, param.PageSize)
    param.Data = arr
    
    OutJson(w, 0, "success", param)
    return
}


