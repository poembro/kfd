package dao

import (
    "database/sql"
    "kfd/internal/logic/db"
    "kfd/internal/logic/model"
    "kfd/pkg/gerrors"
)

type userthirdDao struct{}

var UserthirdDao = new(userthirdDao) 

// Add 插入一条
func (*userthirdDao) Add(item model.Userthird) (int64, error) {
    result, err := db.DBCli.Exec("insert ignore into user_third(user_id,typeid,openid) values(?,?,?)", item.UserId, item.Typeid, item.Openid)
    if err != nil {
        return 0, gerrors.WrapError(err)
    }

    affected, err := result.RowsAffected()
    if err != nil {
        return 0, gerrors.WrapError(err)
    }
    return affected, nil
}

// Get 获取
func (*userthirdDao) Get(Typeid int32, Openid string) (*model.Userthird, error) {
    row := db.DBCli.QueryRow("select user_id,typeid,openid from user_third where typeid = ? and openid = ?", Typeid, Openid)
    item := model.Userthird{
        Typeid: Typeid,
        Openid: Openid,
    }

    err := row.Scan(&item.UserId, &item.Typeid, &item.Openid)
    if err != nil && err != sql.ErrNoRows {
        return nil, gerrors.WrapError(err)
    }

    if err == sql.ErrNoRows {
        return nil, nil
    }

    return &item, err
}
