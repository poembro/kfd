package dao

import (
    "fmt"
    "database/sql"
    "kfd/internal/logic/db"
    "kfd/internal/logic/model"
    "kfd/pkg/gerrors"
)

type userDao struct{}

var (
    USER_TABLE string = "kfd_user" 
    UserDao = new(userDao)
)


func (*userDao) ListUserCount(sex int32) (int) {
    var (
        count int
    )

    row := db.DBCli.QueryRow("SELECT count(id) as count FROM " + USER_TABLE + " where sex = ? limit 1", sex)
    row.Scan(&count)

    return count
}

func (*userDao) ListUser(sex int32, CurrentPage, limit int) ([]model.User, error) {
    sqlstr := "select user_id,nickname,sex,avatar_url,extra,create_time,update_time FROM %s WHERE sex = %d ORDER BY id DESC limit %d, %d"
    sqlstr = fmt.Sprintf(sqlstr, USER_TABLE, sex, (CurrentPage-1) * limit, limit)
    fmt.Println(sqlstr)
    rows, err := db.DBCli.Query(sqlstr)
    if err != nil {
        return nil, gerrors.WrapError(err)
    }

    var users []model.User
    for rows.Next() {
        var item = model.User{ Sex:sex, }
        err := rows.Scan(&item.UserId, &item.Nickname, &item.Sex, &item.AvatarUrl, &item.Extra, &item.CreateTime, &item.UpdateTime)
        if err != nil {
            return nil, gerrors.WrapError(err)
        }
        users = append(users, item)
    }
    return users, nil
}

func (*userDao) AccountByUser(account string) (*model.User, error) {
    row := db.DBCli.QueryRow("select user_id,nickname,sex,avatar_url,extra,create_time,update_time,password from "+ USER_TABLE +" where account = ? limit 1", account)
    
    user := model.User{ Account: account }

    err := row.Scan(&user.UserId, &user.Nickname, &user.Sex, &user.AvatarUrl, &user.Extra, &user.CreateTime, &user.UpdateTime, &user.Password)
    if err != nil && err != sql.ErrNoRows {
        return &user, gerrors.WrapError(err)
    }

    if err == sql.ErrNoRows {
        return &user, err
    }

    return &user, err
}

// Add 插入一条用户信息
func (*userDao) Add(user model.User) (int64, error) {
    fmt.Println(user)
    strsql := fmt.Sprintf("insert ignore into %s(app_id,user_id,nickname,sex,avatar_url,extra,account,password) values(%d,%d,'%s',%d,'%s','%s','%s','%s')", 
        USER_TABLE,user.AppId, user.UserId, user.Nickname, user.Sex, user.AvatarUrl, user.Extra, user.Account, user.Password)
    fmt.Println(strsql)
    
    result, err := db.DBCli.Exec(strsql)
    if err != nil {
        return 0, gerrors.WrapError(err)
    }
    
    affected, err := result.RowsAffected()
    if err != nil {
        return 0, gerrors.WrapError(err)
    }
    return affected, nil
}

// Get 获取用户信息
func (*userDao) Get(appId, userId int64) (*model.User, error) {
    row := db.DBCli.QueryRow("select nickname,sex,avatar_url,extra,create_time,update_time from  "+ USER_TABLE +" where app_id = ? and user_id = ?",
               appId, userId)
    user := model.User{
        AppId:  appId,
        UserId: userId,
    }

    err := row.Scan(&user.Nickname, &user.Sex, &user.AvatarUrl, &user.Extra, &user.CreateTime, &user.UpdateTime)
    if err != nil && err != sql.ErrNoRows {
        return nil, gerrors.WrapError(err)
    }

    if err == sql.ErrNoRows {
        return nil, nil
    }

    return &user, err
}

// Update 更新用户信息
func (*userDao) Update(user model.User) error {
    _, err := db.DBCli.Exec("update  "+ USER_TABLE +" set nickname = ?,sex = ?,avatar_url = ?,extra = ? where app_id = and user_id = ?",
        user.Nickname, user.Sex, user.AvatarUrl, user.Extra, user.AppId, user.UserId)
    if err != nil {
        return gerrors.WrapError(err)
    }

    return nil
}
