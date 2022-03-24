package member

// SignInMemberCmd - 登入會員
type SignInMemberCmd struct {
	// 名稱
	Name string `json:"name"`

	// 密碼
	Password string `json:"password"`
}
