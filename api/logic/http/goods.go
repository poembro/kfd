package http

import (
    "context"
    "net/http"
    "io/ioutil"  
    json "encoding/json" 
    "kfd/internal/logic/model" 
    service "kfd/internal/logic/service"
)

func GoodsInfo(w http.ResponseWriter, r *http.Request) {
    var param = model.Goods{}
    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &param)
    
    if param.Id <= 0 {
        OutJson(w, 1, "goods id error ", param.Id)
        return 
    }
    
    item, err := service.GoodsService.Get(context.TODO(), param)
    if err != nil  {
        OutJson(w, 1, "is not data ", param.Id)
        return
    }
    
    OutJson(w, 0, "success", item)
    return
}

type GoodsPage struct{
    CategoryId     int32  `json:"category_id"`   // sex
    Count          int    `json:"count"`         // count
    TotalPages     int    `json:"totalpages"`    // totalpages
    PageSize       int    `json:"pagesize"`      // pagesize
    CurrentPage    int    `json:"currentpage"`   // currentpage
    Data           []model.Goods `json:"data"`
}

func GoodsList(w http.ResponseWriter, r *http.Request) {
    var (
        param GoodsPage
        opt model.Goods
    )
    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &param)
    
    if  param.PageSize <= 0  {
        param.PageSize = 10
    }
    if  param.CurrentPage <= 0  {
        param.CurrentPage = 1
    }

    opt.CategoryId = param.CategoryId
    count, arr, err := service.GoodsService.ListGoods(context.TODO(), opt, param.CurrentPage, param.PageSize)
    if err != nil  {
        OutJson(w, 1, "is not data ", err)
        return
    }
    param.Count = count
    param.TotalPages = TotalPage(count, param.PageSize)
    param.Data = arr
    
    OutJson(w, 0, "success", param)
    return
}


