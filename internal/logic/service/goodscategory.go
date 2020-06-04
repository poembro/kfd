package service

import (
    "context"
    //"kfd/internal/logic/cache"
    "kfd/internal/logic/dao"
    "kfd/internal/logic/model"
)

type goodsCategoryService struct{}

var GoodsCategoryService = new(goodsCategoryService)

func (*goodsCategoryService) List(ctx context.Context, opt model.GoodsCategory, page, limit int) (count int, arr []model.GoodsCategory, err error) {
    count = dao.GoodsCategoryDao.ListCount(opt)
    
    arr, err = dao.GoodsCategoryDao.List(opt, page, limit)
    if err != nil {
        return 0, nil, err
    } 
     
    return count, arr, err
}

func (*goodsCategoryService) Get(ctx context.Context, opt model.GoodsCategory) (item *model.GoodsCategory, err error) { 
	item, err = dao.GoodsCategoryDao.Get(opt)
	if err != nil {
		return nil, err
	}
    
	return item, err
}
 