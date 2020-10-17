package service

import (
    "fmt"
    "time"
    "context"
    "kfd/pkg/util"
    //"kfd/internal/logic/cache"
    "kfd/internal/logic/dao"
    "kfd/internal/logic/model"
)

type orderService struct{}

var OrderService = new(orderService)

func (*orderService) List(ctx context.Context, opt model.Order, page, limit int) (count int, arr []model.OrderResult, err error) {
    count = dao.OrderDao.ListCount(opt)
    
    arr, err = dao.OrderDao.List(opt, page, limit)
    if err != nil {
        return 0, nil, err
    } 
     
    return count, arr, err
}

func (*orderService) Get(ctx context.Context, opt model.Order) (item *model.OrderResult, err error) { 
	item, err = dao.OrderDao.Get(opt)
	if err != nil {
		return nil, err
	}
    
	return item, err
}


func (*orderService) Save(ctx context.Context, uid int64,opt []map[string]int64) (item *model.Order, err error) { 
    var (
        totalPrice int64
        Order model.Order
        OrderGoods model.OrderGoods
        OrderGoodsArr []model.OrderGoods
    )
    //id生成器获取id
    OrderId, err := util.BuildOrderSn.Get()
    if err != nil {
        return nil, err
    }

    //layout := "Year:2006 Month:01 Day:02 Hour:03 Min:04 Second:05" 
    layout := "20060102"
    dateline := time.Now().Format(layout)
    OrderSn := fmt.Sprintf("O_%s%d",dateline, OrderId) //strconv.Itoa(i)

    user, err := UserService.Get(ctx, int64(1), uid)
    if err != nil  {
        return nil, err
    }
    
    for _, val := range opt {
        goods_id, idok := val["goods_id"]
        goods_num, numok := val["goods_num"]
        if !idok || !numok {
            continue
        }
        goodsInfo, err := GoodsService.Get(ctx, model.Goods{Id:goods_id})
        if err != nil  {
            return nil, err
        }
        
        OrderGoods.OrderSn = OrderSn
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


    Order.OrderSn = OrderSn
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

	return &Order, err
}

