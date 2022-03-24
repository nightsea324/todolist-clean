package member

import "context"

// Usecase -
type Usecase interface {
	// 註冊會員
	RegisterMember(ctx context.Context, cmd *RegisterMemberCmd) (*RegisterMemberEvent, error)
}
