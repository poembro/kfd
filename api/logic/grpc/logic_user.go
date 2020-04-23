package logic

import (
	"context"
	"kfd/pkg/gerrors" 
	"kfd/internal/logic/service"
	"kfd/pkg/pb"
)

type LogicUserServer struct{}

// 登录
func (*LogicUserServer) AccountByUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserResp, error) {
	user, _ := service.UserService.AccountByUser(ctx, req.Account)

	if user == nil {
		return nil, gerrors.ErrUserNotExist
	}

	pbUser := pb.User{
		UserId:     user.UserId,
		Nickname:   user.Nickname,
		Sex:        user.Sex,
		AvatarUrl:  user.AvatarUrl,
		Extra:      user.Extra,
		CreateTime: user.CreateTime.Unix(),
		UpdateTime: user.UpdateTime.Unix(),
		Account :   user.Account,
	}
	return &pb.GetUserResp{User: &pbUser}, nil
}