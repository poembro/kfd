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
    GOODS_key    = "goods:"
    GOODS_expire = 2 * time.Hour
)

type goodsCache struct{}

var GoodsCache = new(goodsCache)

func (*goodsCache) Key(goodsId int64) string {
    return GOODS_key + strconv.FormatInt(goodsId, 10) 
}

// Get 获取
func (c *goodsCache) Get(goodsId int64) (item *model.Goods, err error) {
    err = get(c.Key(goodsId), item)
    if err != nil && err != redis.Nil {
        return nil, gerrors.WrapError(err)
    }
    if err == redis.Nil {
        return nil, nil
    }
    return item, nil
}

// Set 设置
func (c *goodsCache) Set(item model.Goods) error {
    err := set(c.Key(item.Id), item, GOODS_expire)
    if err != nil {
        return gerrors.WrapError(err)
    }
    return nil
}

// Del 删除
func (c *goodsCache) Del(goodsId int64) error {
    _, err := db.RedisCli.Del(c.Key(goodsId)).Result()
    if err != nil {
        return gerrors.WrapError(err)
    }
    return nil
}
