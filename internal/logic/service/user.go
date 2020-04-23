package service

import (
    "context"
    "kfd/internal/logic/cache"
    "kfd/internal/logic/dao"
    "kfd/internal/logic/model"
    "kfd/pkg/gerrors"
)

type userService struct{}

var UserService = new(userService)

func (*userService) GetUserthird(ctx context.Context, Typeid int32, Openid string) (item *model.Userthird, err error) {
    item, err = cache.UserthirdCache.Get(Typeid, Openid)
	if err != nil {
		return nil, err
	}
	if item != nil {
		return item, nil
    }
    
	item, err = dao.UserthirdDao.Get(Typeid, Openid)
	if err != nil {
		return nil, err
	}

	if item != nil {
		err = cache.UserthirdCache.Set(*item)
		if err != nil {
			return nil, err
		}
	}
	return item, nil
}

func (*userService) AddUserthird(ctx context.Context, item model.Userthird) error {
    affected, err := dao.UserthirdDao.Add(item)
    if err != nil {
        return err
    }
    if affected == 0 {
        return gerrors.ErrUserAlreadyExist
    }
    return nil
}


func (*userService) ListUser(ctx context.Context, sex int32, page, limit int) (int, []model.User, error) {
    var (
        count int
        user []model.User
        err error
    )

    count = dao.UserDao.ListUserCount(sex)
   
    user, err = dao.UserDao.ListUser(sex, page, limit)
    if err != nil {
        return 0, nil, err
    } 
     
    return count, user, err
}

//1.添加用户 
func (*userService) Add(ctx context.Context, user model.User) error {
    affected, err := dao.UserDao.Add(user)
    if err != nil {
        return err
    }
    if affected == 0 {
        return gerrors.ErrUserAlreadyExist
    }
    return nil
}

// Get 获取用户信息
func (*userService) Get(ctx context.Context, appId, userId int64) (*model.User, error) {
	user, err := cache.UserCache.Get(appId, userId)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}

	user, err = dao.UserDao.Get(appId, userId)
	if err != nil {
		return nil, err
	}

	if user != nil {
		err = cache.UserCache.Set(*user)
		if err != nil {
			return nil, err
		}
	}
	return user, err
}


func (*userService) Update(ctx context.Context, user model.User) error {
    err := dao.UserDao.Update(user)
    if err != nil {
        return err
    }

    err = cache.UserCache.Del(user.AppId, user.UserId)
    if err != nil {
        return err
    }

    return nil
}



// 通过手机号获取用户信息
func (*userService) AccountByUser(ctx context.Context, account string) (*model.User, error) {
    user, err := dao.UserDao.AccountByUser(account)
    //if err != nil {
    //    return user, err
    //}

    return user, err
}