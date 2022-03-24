package member

import (
	"context"
	"errors"
	"main/service/member/domain/model/aggregate"
	"main/service/member/domain/repo"
	"main/service/member/domain/srv"

	"go.mongodb.org/mongo-driver/bson"
)

// usecase -
type usecase struct {
	memberRepo repo.MemberRepo
	memberSrv  srv.MemberSrv
}

// New -
func New(memberRepo repo.MemberRepo, memberSrv srv.MemberSrv) Usecase {
	return &usecase{
		memberRepo: memberRepo,
		memberSrv:  memberSrv,
	}
}

// RegisterMember - 註冊會員
func (u *usecase) RegisterMember(ctx context.Context, cmd *RegisterMemberCmd) (*RegisterMemberEvent, error) {
	// 轉成model
	member := &aggregate.Member{
		ID:       bson.NewObjectId().Hex(),
		Name:     cmd.Name,
		Password: cmd.Password,
	}

	// 檢查名稱重複
	_, err := u.memberRepo.GetByName(ctx, member.Name)
	if err == nil {
		return nil, errors.New("註冊失敗，名稱重複")
	}
	// 雜湊密碼
	hashedPwd := u.memberSrv.Hash(member.Password)

	member.Password = hashedPwd

	// 建立會員
	if err := u.memberRepo.Create(ctx, member); err != nil {
		return nil, errors.New("註冊失敗，建立錯誤")
	}

	return &RegisterMemberEvent{}, nil
}
