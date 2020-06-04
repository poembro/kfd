package service

import (
    "context"
    //"kfd/internal/logic/cache"
    "kfd/internal/logic/dao"
    "kfd/internal/logic/model"
)

type orderService struct{}

var OrderService = new(orderService)

func (*orderService) List(ctx context.Context, opt model.Order, page, limit int) (count int, arr []model.OrderResult, err error) {
    count = dao.OrdersDao.ListCount(opt)
    
    arr, err = dao.OrdersDao.List(opt, page, limit)
    if err != nil {
        return 0, nil, err
    } 
     
    return count, arr, err
}

func (*orderService) Get(ctx context.Context, opt model.Order) (item *model.OrderResult, err error) { 
	item, err = dao.OrdersDao.Get(opt)
	if err != nil {
		return nil, err
	}
    
	return item, err
}
 

func (*orderService) Get(ctx context.Context, opt model.Order) (item *model.OrderResult, err error) { 
	item, err = dao.OrdersDao.Get(opt)
	if err != nil {
		return nil, err
	}
    
	return item, err
}


func (*orderService) Save(ctx context.Context, uid int64,opt []map[string]int64) (item *model.OrderResult, err error) { 
    var (
        totalPrice int64
        Order model.Order
        OrderGoods model.OrderGoods
        OrderGoodsArr []model.OrderGoods
    )
    //id生成器获取id
    OrderSn, err := util.OrderBusinessId.Get()
    if err != nil {
        OutJson(w, 1, "id error", nil)
        return
    }

    user, err := UserService.Get(ctx, int64(1), uid)
    if err != nil  {
        OutJson(w, 1, "error user not register", uid)
        return
    }
    
    for _, val := range opt {
        goods_id, idok := val["goods_id"]
        goods_num, numok := val["goods_num"]
        if !idok || !numok {
            continue
        } 
        goodsInfo, err := GoodsService.Get(ctx, model.Goods{Id:goods_id})
        if err != nil  {
            OutJson(w, 1, "error item not ", param.Id)
            return
        }
        
        OrderGoods.OrderSn = string(OrderSn)
        OrderGoods.GoodsId = goods_id
        OrderGoods.GoodsNum = goods_num 
        OrderGoods.GoodsName = goodsInfo.Name 
        OrderGoods.Thumb = goodsInfo.Thumb 
        OrderGoods.GoodsPrice = goodsInfo.GoodsPrice 

        OrderGoodsArr = append(OrderGoodsArr, OrderGoods)
        //写入
        totalPrice += goods_num * goodsInfo.GoodsPrice
    }

    dao.OrderDao.AddOrderGoods(OrderGoodsArr)


    Order.OrderSn = string(OrderSn)
    Order.Uid = uid
    Order.Status = 1
    Order.Payid = 1
    Order.Ispay = 1 
    Order.Realname = user.Nickname
    Order.Mobile = user.Account 
    Order.Address = user.Extra
    Order.TotalPrice = totalPrice
    //写入
    dao.OrderDao.AddOrder(Order)

	return item, err
}

