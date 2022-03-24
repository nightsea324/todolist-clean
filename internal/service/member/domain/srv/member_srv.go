package srv

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"main/common/auth"
	"main/service/member/domain/model/aggregate"
	"strings"
)

// Hash - 雜湊密碼
func (m *memberSrv) Hash(key1 string) string {
	h := hmac.New(sha256.New, []byte(HASHKEY))
	data := strings.TrimSpace(key1) + "/nightsea/" + strings.TrimSpace(HASHKEY)

	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// CreateToken - 建立權杖
func (m *memberSrv) CreateToken(ctx context.Context, member *aggregate.Member) (string, error) {
	memberToken := &auth.MemberToken{
		ID:   member.ID,
		Name: member.Name,
	}
}
