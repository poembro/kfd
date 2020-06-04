package model

import (
    "time"
)

// goods
type Goods struct {
    Id           int64     `json:"id"`            // 用户id
    Name         string    `json:"name"`         // 商品名称
    Description  string    `json:"description"`   // 描述
    CategoryId   int32     `json:"category_id"`   // 分类id
    IsOnSale     int32     `json:"is_on_sale"`    // 上架状态 1-上架 0-下架
    Thumb        string    `json:"thumb"`         // 商品缩略图
    GoodsPrice   float64    `json:"goods_price"`   // 商品价格
    MarketPrice  float64    `json:"market_price"`  // 市场价格
    SellsNum     int64    `json:"sells_num"`     // 销售数 
    CommentsNum  int64    `json:"comments_num"`  // 评价数
    FavoritesNum int64    `json:"favorites_num"` // 收藏数
    IsHot        int32    `json:"is_hot"`        // 推荐，首页显示 
    GoodsSort    int64    `json:"goods_sort"`    // 排序字段
    Inventory    int64    `json:"inventory"`     // 库存 
    Visible      int32    `json:"visible"`       // 附加属性  
    CreateTime time.Time  `json:"-"`             // 创建时间
    UpdateTime time.Time  `json:"-"`             // 更新时间 
}
