package dao

import (
    "fmt"
    "strings"
    "database/sql"
    "kfd/internal/logic/db"
    "kfd/internal/logic/model"
    "kfd/pkg/gerrors"
)


type goodsCategoryDao struct{}

var (
    GOODSCATEGORY_TABLE string = "kfd_goods_category"
    GoodsCategoryDao = new(goodsCategoryDao)
)

func (d *goodsCategoryDao) ListCount(opt model.GoodsCategory) (int) {
    var (
        count int
        sqlstr string
    )
    sqlstr = d._handle(opt)
    sqlstr = "SELECT count(id) as count FROM " + GOODSCATEGORY_TABLE + " WHERE " + sqlstr + " limit 1" 
    fmt.Println("sql:", sqlstr)
    row := db.DBCli.QueryRow(sqlstr)
    row.Scan(&count)

    return count
}
 
func (d *goodsCategoryDao) List(opt model.GoodsCategory, CurrentPage, limit int) ([]model.GoodsCategory, error) {
    var (
        sqlstr string
        item model.GoodsCategory
    )
    sqlstr = d._handle(opt)
    sqlstr = "select id,name,sort_id,visible from " + GOODSCATEGORY_TABLE + " WHERE " + sqlstr + " ORDER BY id DESC limit ?, ?" 
    fmt.Println("sql:", sqlstr)
    rows, err := db.DBCli.Query(sqlstr, (CurrentPage-1) * limit, limit)
    if err != nil {
        return nil, gerrors.WrapError(err)
    }
    
    arr := make([]model.GoodsCategory, 0, 5)
    for rows.Next() {
        err := rows.Scan(
             &item.Id, &item.Name, 
             &item.SortId,
             &item.Visible )
        if err != nil {
            return nil, gerrors.WrapError(err)
        }
        arr = append(arr, item)
    }
    return arr, nil
}


func (d *goodsCategoryDao) Get(opt model.GoodsCategory) (*model.GoodsCategory, error) {
    var ( 
        sqlstr string
        item model.GoodsCategory
    )
    sqlstr = d._handle(opt)
    sqlstr = "SELECT id,name,sort_id,visible FROM " + GOODSCATEGORY_TABLE + " WHERE " + sqlstr + " limit 1" 
    fmt.Println("sql:", sqlstr) 

    row := db.DBCli.QueryRow(sqlstr)
    err := row.Scan(
        &item.Id, 
        &item.Name, 
        &item.SortId,
        &item.Visible )
    fmt.Println("item:", item)
    if err != nil && err != sql.ErrNoRows {
        return nil, gerrors.WrapError(err)
    }

    if err == sql.ErrNoRows {
        return nil, nil
    }
    
    return &item, err
}


func (d *goodsCategoryDao) _handle(opt model.GoodsCategory) (sqlstr string){
    var (
        where []string
    )

    if opt.Id > 0 {
        tmp := fmt.Sprintf("id=%d", opt.Id)
        where = append(where, tmp)
    }

    if opt.Name != "" {
        tmp := fmt.Sprintf("name LIKE \"%%%s\" ", opt.Name)
        where = append(where, tmp)
    }
    
    where = append(where, "1=1")
	sqlstr = strings.Join(where, " AND ")
    return sqlstr
}