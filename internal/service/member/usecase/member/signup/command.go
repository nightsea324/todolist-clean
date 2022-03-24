package member

// RegisterMemberCmd - 註冊會員
type RegisterMemberCmd struct {
	// 名稱
	Name string `json:"name"`

	// 密碼
	Password string `json:"password"`
}
