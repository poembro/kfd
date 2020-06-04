package model

import (
    "time"
)

// goods
type GoodsCategory struct {
    Id           int64     `json:"id"`            // 用户id
    Name         string    `json:"name"`         // 名称
    SortId       int64    `json:"sort_id"`    // 排序字段 
    Visible      int32    `json:"visible"`       // 附加属性  
    CreateTime time.Time  `json:"-"`             // 创建时间
    UpdateTime time.Time  `json:"-"`             // 更新时间 
}
