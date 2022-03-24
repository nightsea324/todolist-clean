package member

import "context"

// Usecase -
type Usecase interface {
	// 登入會員
	SignInMember(ctx context.Context, cmd *SignInMemberCmd) (*SignInMemberEvent, error)
}
