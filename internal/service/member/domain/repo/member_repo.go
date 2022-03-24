package repo

import (
	"context"
	"main/service/member/domain/model/aggregate"
)

// MemberRepo -
type MemberRepo interface {
	// Create - 建立會員
	Create(ctx context.Context, member *aggregate.Member) error

	// GetByName - 透過名稱取得會員
	GetByName(ctx context.Context, name string) (*aggregate.Member, error)
}
