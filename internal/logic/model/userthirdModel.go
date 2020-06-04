package model


// User 账户
type Userthird struct { 
    UserId     int64     `json:"user_id"`  // 手机号
    Typeid     int32     `json:"typeid"`   // 第三方类型 1-Facebook
    Openid     string    `json:"openid"`   // 昵称 
}
 