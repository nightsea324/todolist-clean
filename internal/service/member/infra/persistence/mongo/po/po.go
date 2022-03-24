package po

import "time"

// MemberPo -
type MemberPo struct {
	// 會員編號
	ID string `bson:"id"`
	// 名稱
	Name string `bson:"name"`
	// 密碼
	Password string `bson:"password"`
	// 建立時間
	CreatedAt time.Time `bson:"created_at"`
}
