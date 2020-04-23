package cache

import (
    "kfd/internal/logic/db"
    "kfd/internal/logic/model"
    "kfd/pkg/gerrors"
    "strconv"  //strconv包对字符串转化成其他基本类型的数据进行判断
    "time"

    "github.com/go-redis/redis"
)

const ( 
    USERTHIRD_key    = "userthird:"
    USERTHIRD_expire = 2 * time.Hour
)

type userthirdCache struct{}

var UserthirdCache = new(userthirdCache)

func (*userthirdCache) Key(Typeid int32, Openid string) string {
    return USERTHIRD_key + strconv.FormatInt(int64(Typeid), 10) + ":" + Openid
}

// Get 获取用户缓存
func (c *userthirdCache) Get(Typeid int32, Openid string) (*model.Userthird, error) {
    var item model.Userthird
    err := get(c.Key(Typeid, Openid), &item)
    if err != nil && err != redis.Nil {
        return nil, gerrors.WrapError(err)
    }
    if err == redis.Nil {
        return nil, nil
    }
    return &item, nil
}

// Set 设置用户缓存
func (c *userthirdCache) Set(item model.Userthird) error {
    err := set(c.Key(item.Typeid, item.Openid), item, USERTHIRD_expire)
    if err != nil {
        return gerrors.WrapError(err)
    }
    return nil
}

// Del 删除用户缓存
func (c *userthirdCache) Del(Typeid int32, Openid string) error {
    _, err := db.RedisCli.Del(c.Key(Typeid, Openid)).Result()
    if err != nil {
        return gerrors.WrapError(err)
    }
    return nil
}
