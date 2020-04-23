package service

import (
    "context"
    //"log"
    //"kfd/internal/logic/dao"
    "kfd/pkg/util"
    "time"
)

type authService struct{}

var AuthService = new(authService)
 
// 构建Token
func (*authService) BuildToken(ctx context.Context, userId int64, account string) (token string, err error) {
    expire := time.Now().Add(24 * 7 * time.Hour).Unix()
    token, err = util.GetToken(1, userId, account, expire, util.PublicKey)
    return token, err
}

// VerifySecretKey 对用户秘钥进行校验
func (*authService) VerifyToken(ctx context.Context, token string) (*util.TokenInfo, error) { 
    info, err := util.DecryptToken(token, util.PrivateKey)
    if err != nil {
        return nil, err
    }
    //log.Println("打印 秘钥进行校验 参数：",info)
    if info.Account == "" {
        return nil, err
    }
    
    if info.Expire < time.Now().Unix() {
        return nil, err
    }
    
    return info, err
}
