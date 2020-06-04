package dao

import (
    "fmt"
    "strings"
    "database/sql"
    "kfd/internal/logic/db"
    "kfd/internal/logic/model"
    "kfd/pkg/gerrors"
)


type goodsDao struct{}

var (
    GOODS_TABLE string = "kfd_goods"
    GoodsDao = new(goodsDao)
)

func (d *goodsDao) ListGoodsCount(opt model.Goods) (int) {
    var (
        count int
        sqlstr string
    )
    sqlstr = d._handle(opt)
    sqlstr = "SELECT count(id) as count FROM " + GOODS_TABLE + " WHERE " + sqlstr + " limit 1" 
    fmt.Println("sql:", sqlstr)
    row := db.DBCli.QueryRow(sqlstr)
    row.Scan(&count)

    return count
}

func (d *goodsDao) ListGoods(opt model.Goods, CurrentPage, limit int) ([]model.Goods, error) {
    var (
        sqlstr string
        item model.Goods
    )
    sqlstr = d._handle(opt)
    sqlstr = "select id,name,description,category_id,is_on_sale,thumb,goods_price,market_price,sells_num,comments_num,favorites_num,is_hot,goods_sort,inventory,visible from " + GOODS_TABLE + " WHERE " + sqlstr + " ORDER BY id DESC limit ?, ?" 
    fmt.Println("sql:", sqlstr)
    rows, err := db.DBCli.Query(sqlstr, (CurrentPage-1) * limit, limit)
    if err != nil {
        return nil, gerrors.WrapError(err)
    }
    
    arr := make([]model.Goods, 0, 5)
    for rows.Next() {
        err := rows.Scan(
             &item.Id, &item.Name, 
             &item.Description,
             &item.CategoryId,
             &item.IsOnSale, 
             &item.Thumb, &item.GoodsPrice,  
             &item.MarketPrice, &item.SellsNum, 
             &item.CommentsNum, &item.FavoritesNum,  
             &item.IsHot, &item.GoodsSort,
             &item.Inventory, &item.Visible )
        if err != nil {
            return nil, gerrors.WrapError(err)
        }
        arr = append(arr, item)
    }
    return arr, nil
}


func (d *goodsDao) Get(opt model.Goods) (*model.Goods, error) {
    var ( 
        sqlstr string
        item model.Goods
    )
    sqlstr = d._handle(opt)
    sqlstr = "SELECT name,description,category_id,is_on_sale,thumb,goods_price,market_price,sells_num,comments_num,favorites_num,is_hot,goods_sort,inventory,visible FROM " + GOODS_TABLE + " WHERE " + sqlstr + " limit 1" 
    fmt.Println("sql:", sqlstr) 

    row := db.DBCli.QueryRow(sqlstr)
    err := row.Scan(&item.Name, &item.Description, 
        &item.CategoryId, &item.IsOnSale, 
        &item.Thumb, &item.GoodsPrice,
        &item.MarketPrice, &item.SellsNum,
        &item.CommentsNum, &item.FavoritesNum,
        &item.IsHot, &item.GoodsSort,
        &item.Inventory, &item.Visible)
    
    if err != nil && err != sql.ErrNoRows {
        return nil, gerrors.WrapError(err)
    }

    if err == sql.ErrNoRows {
        return nil, nil
    }

    return &item, err
}


func (d *goodsDao) _handle(opt model.Goods) (sqlstr string){
    var (
        where []string
    )

    if opt.Id > 0 {
        tmp := fmt.Sprintf("id=%d", opt.Id)
        where = append(where, tmp)
    }

    if opt.CategoryId > 0 {
        tmp := fmt.Sprintf("category_id=%d", opt.CategoryId)
        where = append(where, tmp)
    }
    

    if opt.Name != "" {
        tmp := fmt.Sprintf("name LIKE \"%%%s\" ", opt.Name)
        where = append(where, tmp)
    }
    
    where = append(where, "1=1")
	sqlstr = strings.Join(where, " and ")
    return sqlstr
}