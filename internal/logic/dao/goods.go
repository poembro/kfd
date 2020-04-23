package dao

import (
    "database/sql"
    "kfd/internal/logic/db"
    "kfd/internal/logic/model"
    "kfd/pkg/gerrors"
)

type goodsDao struct{}

var GoodsDao = new(goodsDao)

func (*goodsDao) ListGoodsCount(category_id int32) (int) {
    var (
        count int
    )

    row := db.DBCli.QueryRow("SELECT count(id) as count FROM goods where category_id = ? limit 1", category_id)
    row.Scan(&count)

    return count
}

func (*goodsDao) ListGoods(category_id int32, CurrentPage, limit int) ([]model.Goods, error) {
    rows, err := db.DBCli.Query(`
        select id,name,description,is_on_sale,thumb,goods_price,market_price,sells_num,comments_num,favorites_num,is_hot,goods_sort,inventory,visible
        from goods
        where category_id = ? ORDER BY id DESC limit ?, ?`, category_id, (CurrentPage-1) * limit, limit)
    if err != nil {
        return nil, gerrors.WrapError(err)
    }

    arr := make([]model.Goods, 0, 5)
    for rows.Next() {
        var item = model.Goods{ CategoryId:category_id, }
        err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.IsOnSale, 
             &item.Thumb, &item.GoodsPrice,  &item.MarketPrice, &item.SellsNum, 
             &item.CommentsNum, &item.FavoritesNum,  &item.IsHot, &item.GoodsSort,
             &item.Inventory, &item.Visible )
        if err != nil {
            return nil, gerrors.WrapError(err)
        }
        arr = append(arr, item)
    }
    return arr, nil
}

func (*goodsDao) Get(goodsId int64) (*model.Goods, error) {
    item := model.Goods{Id: goodsId,}

    row := db.DBCli.QueryRow("select name,description,category_id,is_on_sale,thumb,goods_price,market_price,sells_num,comments_num,favorites_num,is_hot,goods_sort,inventory,visible from goods where id = ?", goodsId)

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

