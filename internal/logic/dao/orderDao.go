package dao

import (
    "fmt"
    "time"
    "strings"
    "database/sql"
    "kfd/internal/logic/db"
    "kfd/internal/logic/model"
    "kfd/pkg/gerrors"
)

type orderDao struct{}

var (
    ORDER_TABLE string = "kfd_order"
    ORDERGOODS_TABLE string = "kfd_order_goods"
    OrdersDao = new(orderDao)
)

func (d *orderDao) ListCount(opt model.Order) (int) {
    var (
        count int
        sqlstr string
    )
    sqlstr = d._handle(opt)
    sqlstr = "SELECT count(id) as count FROM " + ORDER_TABLE + " WHERE " + sqlstr + " limit 1" 
    fmt.Println("sql:", sqlstr)
    row := db.DBCli.QueryRow(sqlstr)
    row.Scan(&count)

    return count
}

func (d *orderDao) List(opt model.Order, CurrentPage, limit int) ([]model.OrderResult, error) {
    var (
        sqlstr string
        oinfo model.Order 
    )
    sqlstr = d._handle(opt)
    sqlstr = "select order_sn,uid,status,payid,paytime,ispay,shiptime,finished_time,refund_msg,cancel_type,cancel_remark,realname,mobile,shopname,address,total_price,remark,visible,create_time,update_time from " + ORDER_TABLE + " WHERE " + sqlstr + " ORDER BY id DESC limit ?, ?" 
    fmt.Println("sql:", sqlstr)
    rows, err := db.DBCli.Query(sqlstr, (CurrentPage-1) * limit, limit)
    if err != nil {
        return nil, gerrors.WrapError(err)
    }
    arr := make([]model.Order, 0, 5)
    for rows.Next() {
        err := rows.Scan(&oinfo.OrderSn, 
            &oinfo.Uid, &oinfo.Status,
            &oinfo.Payid, &oinfo.Paytime,
            &oinfo.Ispay, &oinfo.Shiptime,
            &oinfo.FinishedTime, &oinfo.RefundMsg,  
            &oinfo.CancelType,  &oinfo.CancelRemark,
            &oinfo.Realname, &oinfo.Mobile, 
            &oinfo.Shopname,  &oinfo.Address, 
            &oinfo.TotalPrice, &oinfo.Remark, 
            &oinfo.Visible,
            &oinfo.CreateTime,
            &oinfo.UpdateTime)
        if err != nil {
            return nil, gerrors.WrapError(err)
        }

        tt , _:= time.Parse("2006-01-02T15:04:05Z07:00", oinfo.CreateTime)
        oinfo.CreateTime = tt.Format("2006-01-02 15:04:05") 
        tt1 , _:= time.Parse("2006-01-02T15:04:05Z07:00", oinfo.UpdateTime)
        oinfo.UpdateTime = tt1.Format("2006-01-02 15:04:05")

        arr = append(arr, oinfo)
    }
    
    result := make([]model.OrderResult, 0, 10) 
    for _, val := range arr {
        olist, err := d._GetOrderGoods(val.OrderSn) //去取对应 ordersn的所有商品
        if err != nil {
            return nil, gerrors.WrapError(err)
        }
        tmp := model.OrderResult{Order:val, OrderGoods:olist} 
        result = append(result, tmp)
    }

    return result, nil
}


func (d *orderDao) Get(opt model.Order) (*model.OrderResult, error) {
    var ( 
        sqlstr string
        oinfo model.Order
        result model.OrderResult
    )
    sqlstr = d._handle(opt)
    sqlstr = "SELECT order_sn,uid,status,payid,paytime,ispay,shiptime,finished_time,refund_msg,cancel_type,cancel_remark,realname,mobile,shopname,address,total_price,remark,visible,create_time,update_time FROM " + ORDER_TABLE + " WHERE " + sqlstr + " limit 1" 
    fmt.Println("sql:", sqlstr) 
    
    row := db.DBCli.QueryRow(sqlstr)
    err := row.Scan(&oinfo.OrderSn, 
        &oinfo.Uid, 
        &oinfo.Status,
        &oinfo.Payid, 
        &oinfo.Paytime,
        &oinfo.Ispay,
        &oinfo.Shiptime,
        &oinfo.FinishedTime,
        &oinfo.RefundMsg,
        &oinfo.CancelType,
        &oinfo.CancelRemark,
        &oinfo.Realname,
        &oinfo.Mobile, 
        &oinfo.Shopname, 
        &oinfo.Address, 
        &oinfo.TotalPrice,  
        &oinfo.Remark, 
        &oinfo.Visible,
        &oinfo.CreateTime,
        &oinfo.UpdateTime) 
    if err != nil  {
        return nil, gerrors.WrapError(err)
    }
     
    if err == sql.ErrNoRows { //结果集中没有数据行
        return nil, nil
    }
    //时间格式转换一下
    tt , _:= time.Parse("2006-01-02T15:04:05Z07:00", oinfo.CreateTime)
    oinfo.CreateTime = tt.Format("2006-01-02 15:04:05") 
    tt1 , _:= time.Parse("2006-01-02T15:04:05Z07:00", oinfo.UpdateTime)
    oinfo.UpdateTime = tt1.Format("2006-01-02 15:04:05")
    
    arr, err := d._GetOrderGoods(opt.OrderSn)
    result = model.OrderResult{Order:oinfo, OrderGoods:arr}
    return &result, err
}

func (d *orderDao) _GetOrderGoods(OrderSn string) ([]model.OrderGoods, error) {
    var ( 
        sqlstr string
        olist model.OrderGoods
    )
    sqlstr = "select id,order_sn,goods_id,refund_msg,cancel_type,cancel_remark,goods_name,thumb,goods_num,goods_price,goods_spec,status,visible,create_time,update_time from " + ORDERGOODS_TABLE + " WHERE order_sn='" + OrderSn + "' ORDER BY id DESC"
    fmt.Println("sql:", sqlstr)
    rows, err := db.DBCli.Query(sqlstr)
    if err != nil {
        return nil, gerrors.WrapError(err)
    }
    
    arr := make([]model.OrderGoods, 0, 5)
    for rows.Next() {
        err := rows.Scan(
            &olist.Id, 
            &olist.OrderSn, 
            &olist.GoodsId, 
            &olist.RefundMsg, 
            &olist.CancelType, 
            &olist.CancelRemark, 
            &olist.GoodsName, 
            &olist.Thumb, 
            &olist.GoodsNum, 
            &olist.GoodsPrice, 
            &olist.GoodsSpec, 
            &olist.Status, 
            &olist.Visible,
            &olist.CreateTime,
            &olist.UpdateTime)
        if err != nil {
            return nil, gerrors.WrapError(err)
        }
        tt , _:= time.Parse("2006-01-02T15:04:05Z07:00", olist.CreateTime)
        olist.CreateTime = tt.Format("2006-01-02 15:04:05") 
        tt1 , _:= time.Parse("2006-01-02T15:04:05Z07:00", olist.UpdateTime)
        olist.UpdateTime = tt1.Format("2006-01-02 15:04:05")

        arr = append(arr, olist)
    }
    return arr, err
}





func (d *orderDao) _handle(opt model.Order) (sqlstr string){
    var (
        tmp string
        where []string
    )

    if opt.Id > 0 {
        tmp = fmt.Sprintf("id=%d", opt.Id)
        where = append(where, tmp)
    }

    if opt.Uid > 0 {
        tmp = fmt.Sprintf("uid=%d", opt.Uid)
        where = append(where, tmp)
    }

    if opt.Status >= 0 {
        tmp = fmt.Sprintf("status=%d", opt.Status)
        where = append(where, tmp)
    }


    if opt.OrderSn != "" {
        tmp = fmt.Sprintf("order_sn='%s'", opt.OrderSn)
        where = append(where, tmp)
    }
    
    where = append(where, "1=1")
	sqlstr = strings.Join(where, " and ")
    return sqlstr
}




// Add 插入一条用户信息
func (d *orderDao) AddOrderGoods(opt []model.OrderGoods) (int64, error) {
    result, err := db.DBCli.Exec(
        "insert ignore into  "+ USER_TABLE +"(app_id,user_id,nickname,sex,avatar_url,extra,account,password) values(?,?,?,?,?,?,?,?),(?,?,?,?,?,?,?,?)",
        user.AppId, user.UserId, user.Nickname, user.Sex, user.AvatarUrl, user.Extra, user.Account, user.Password)

    if err != nil {
        return 0, gerrors.WrapError(err)
    }

    affected, err := result.RowsAffected()
    if err != nil {
        return 0, gerrors.WrapError(err)
    }
    return affected, nil
}


// Add 插入一条用户信息
func (d *orderDao) AddOrder(opt model.Order) (int64, error) {
    result, err := db.DBCli.Exec("insert ignore into  "+ USER_TABLE +"(app_id,user_id,nickname,sex,avatar_url,extra,account,password) values(?,?,?,?,?,?,?,?)",
        user.AppId, user.UserId, user.Nickname, user.Sex, user.AvatarUrl, user.Extra, user.Account, user.Password)
    if err != nil {
        return 0, gerrors.WrapError(err)
    }

    affected, err := result.RowsAffected()
    if err != nil {
        return 0, gerrors.WrapError(err)
    }
    return affected, nil
}


