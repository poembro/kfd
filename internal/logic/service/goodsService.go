package service

import (
    "context"
    "kfd/internal/logic/cache"
    "kfd/internal/logic/dao"
    "kfd/internal/logic/model"
)

type goodsService struct{}

var GoodsService = new(goodsService)

func (*goodsService) ListGoods(ctx context.Context, opt model.Goods, page, limit int) (count int, arr []model.Goods, err error) {
    count = dao.GoodsDao.ListGoodsCount(opt)
    
    arr, err = dao.GoodsDao.ListGoods(opt, page, limit)
    if err != nil {
        return 0, nil, err
    } 
     
    return count, arr, err
}

func (*goodsService) Get(ctx context.Context, opt model.Goods) (item *model.Goods, err error) { 
	item, err = dao.GoodsDao.Get(opt)
	if err != nil {
		return nil, err
	}
    
	return item, err
}


func (*goodsService) GetCache(ctx context.Context, opt model.Goods) (item *model.Goods, err error) {
	item, err = cache.GoodsCache.Get(opt.Id)
	if err != nil {
		return nil, err
	}
	if item != nil {
		return item, nil
	}

	item, err = dao.GoodsDao.Get(opt)
	if err != nil {
		return nil, err
	}

	if item != nil {
		err = cache.GoodsCache.Set(*item)
		if err != nil {
			return nil, err
		}
	}
	return item, err
}
