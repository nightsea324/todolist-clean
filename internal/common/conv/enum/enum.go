package enum

// TodoListStatus - 待辦事項狀態
type TodoListStatus uint32

const (
	TodoListStatusNone TodoListStatus = iota

	// 1 已完成
	TodoListStatusDone

	// 2 正在完成
	TodoListStatusInProgress

	// 3 待完成
	TodoListStatusToDo

	// 4 已取消
	TodoListStatusCancel
)
