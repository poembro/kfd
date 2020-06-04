package model

type OrderResult struct{
    Order Order  
    OrderGoods []OrderGoods
}

// Order
type Order struct {
    Id           int64     `json:"id"` // 用户id
    OrderSn      string    `json:"order_sn"` // 订单号
    Uid          int64    `json:"uid"`
    Status       int32    `json:"status"`  //订单状态 1-待付款 2-已付款 3-待配送  4-配送中  5-已完成 6-已取消
    Payid        int32    `json:"payid"`   //付款方式 1-建行
    Paytime      int64    `json:"paytime"`  //付款时间
    Ispay        int32    `json:"ispay"`  //是否付款 1-是  2-否 
    Shiptime     int64    `json:"shiptime"` //发货时间
    FinishedTime int64   `json:"finished_time"` //订单完成时间
    RefundMsg  string    `json:"refund_msg"` //退款提示 
    CancelType int32     `json:"cancel_type"`  //取消原因 1-手机号码不是本人 2-用户要求取消 3-其它
    CancelRemark string  `json:"cancel_remark"` //取消备注 
    Realname   string    `json:"realname"` //姓名
    Mobile     string    `json:"mobile"` //手机号码
    Shopname   string    `json:"shopname"` //店铺
    Address    string    `json:"address"` //地址
    TotalPrice int64     `json:"total_price"` //总价格 
    Remark     string    `json:"remark"` //买家留言 
    Visible      int32    `json:"visible"`       // 附加属性  
    CreateTime string  `json:"create_time,omitempty"` // omitempty 当该字段的值为该字段类型的零值时，忽略该字段
    UpdateTime string `json:"update_time,omitempty"`
}



// OrderGoods 
type OrderGoods struct {
    Id            int64     `json:"id"` // 用户id
    OrderSn       string    `json:"order_sn"` // 订单号
    GoodsId       int64    `json:"goods_id"` 
    RefundMsg     string    `json:"refund_msg"` //退款提示 
    CancelType    int32     `json:"cancel_type"`  //取消原因 1-手机号码不是本人 2-用户要求取消 3-其它
    CancelRemark  string  `json:"cancel_remark"` //取消备注  
    GoodsName     string    `json:"goods_name"` // 
    Thumb         string    `json:"thumb"` // 缩略图
    GoodsNum      int64     `json:"goods_num"`   //数量 
    GoodsPrice    int64    `json:"goods_price"`
    GoodsSpec     int64    `json:"goods_spec"` //规格
    Status        int32    `json:"status"`  //订单状态 0-待付款 1-已付款 2-待配送  3-配送中  4-已完成 5-已取消
    Visible       int32    `json:"visible"`       // 附加属性  
    CreateTime    string  `json:"create_time"`             // 创建时间
    UpdateTime    string  `json:"update_time"`             // 更新时间 
}
