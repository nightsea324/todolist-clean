package srv

// HASHKEY -
const HASHKEY = "nightsea-todolist-2022"

type memberSrv struct {
}

// New -
func New() MemberSrv {
	return &memberSrv{}
}
