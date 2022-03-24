package srv

import (
	"context"
	"main/service/member/domain/model/aggregate"
)

// MemberSrv -
type MemberSrv interface {
	// Hash - 雜湊密碼
	Hash(key1 string) string

	// CreateToken - 建立權杖
	CreateToken(ctx context.Context, member *aggregate.Member) (string, error)
}
