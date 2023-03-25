package constant

const (
	UserTableName        = "user"
	TodoTableName        = "todo"
	DefaultAvatarAddress = "https://www.baidu.com/avatar/default-avatar.jpg"
	CodePrefix           = "todo-list-backup-code-%s"
	CodeLength           = 6
	CodeExpires          = 60 // seconds
	TokenOfHeaderKey     = "Authorization"
	IDOfContextKey       = "id"
	UsernameOfContextKey = "username"
	EmailOfSubject       = "Todo List Backup"
)
