package http

import (
    "context"
    "net/http"
    "io/ioutil"  
    json "encoding/json" 
    "kfd/internal/logic/model" 
    service "kfd/internal/logic/service"
)


//接收参数格式: [{"goods_id":33,"goods_num":2},{"goods_id":33,"goods_num":1},]
func OrderSave(w http.ResponseWriter, r *http.Request) {
    var (
        user_id int64
        ok bool
        param []map[string]int64
    )
    
    user_id, ok = r.Context().Value("user_id").(int64); //作为赋值中的类型字符串:需要类型断言
    if user_id == 0 || !ok {
        OutJson(w, 1, "token error r.Context().Value ", user_id)
        return  
    }
    
    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &param)

    if len(param) == 0 {
        OutJson(w, 1, "error param is null ", param)
        return
    }
    
    item, err := service.OrderService.Save(context.TODO(), user_id, param)
    if err != nil  {
        OutJson(w, 1, "not data ", item.OrderSn)
        return
    }
    
    OutJson(w, 0, "success", item)
    return
}


func OrderInfo(w http.ResponseWriter, r *http.Request) {
    var (
        user_id int64
        ok bool
        param model.Order 
        err error
        item *model.OrderResult
    )

    user_id, ok = r.Context().Value("user_id").(int64); //作为赋值中的类型字符串:需要类型断言
    if user_id == 0 || !ok {
        OutJson(w, 1, "token error1 r.Context().Value ", user_id)
        return  
    }
    
    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &param)
    if param.OrderSn == "" {
        OutJson(w, 1, "order id error ", param.OrderSn)
        return 
    }
    
    param.Uid = user_id
    item, err = service.OrderService.Get(context.TODO(), param)
    if err != nil  {
        OutJson(w, 1, "not data ", param.OrderSn)
        return
    }
    
    OutJson(w, 0, "success", item)
    return
}

type OrderPage struct{
    Uid            int64    `json:"uid"`
    Status         int32  `json:"status"`  //订单状态 0-待付款 1-已付款 2-待配送  3-配送中  4-已完成 5-已取消
    Count          int    `json:"count"`        // count
    TotalPages     int    `json:"totalpages"`   // totalpages
    PageSize       int    `json:"pagesize"`     // pagesize
    CurrentPage    int    `json:"currentpage"`  // currentpage
    Data           []model.OrderResult `json:"data"`
}

func OrderList(w http.ResponseWriter, r *http.Request) {
    var (
        user_id int64
        ok bool
        param OrderPage
        opt model.Order
    )
    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &param)
    
    user_id, ok = r.Context().Value("user_id").(int64); //作为赋值中的类型字符串:需要类型断言
    if user_id == 0 || !ok {
        OutJson(w, 1, "token error1 r.Context().Value ", user_id)
        return  
    }

    if  param.PageSize <= 0  {
        param.PageSize = 10
    }
    if  param.CurrentPage <= 0  {
        param.CurrentPage = 1
    }

    if  param.Status <= 0  {
        param.Status = 0
    }
    
    opt.Uid = user_id
    opt.Status = param.Status
    count, arr, err := service.OrderService.List(context.TODO(),  opt, param.CurrentPage, param.PageSize)
    if err != nil  {
        OutJson(w, 1, "not data ", err)
        return
    }
    param.Count = count
    param.TotalPages = TotalPage(count, param.PageSize)
    param.Data = arr
    
    OutJson(w, 0, "success", param)
    return
}
 