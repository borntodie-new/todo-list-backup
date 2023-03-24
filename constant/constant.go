package constant

const (
	UserTableName        = "user"
	TodoTableName        = "todo"
	DefaultAvatarAddress = "https://www.baidu.com/avatar/default-avatar.jpg"
	DefaultLimit         = 10
	DefaultOffset        = 0
	CodePrefix           = "todo-list-backup-code-%s"
	CodeLength           = 6
	CodeExpires          = 60 // seconds
	TokenExpires         = 24 // hours
	TokenOfHeaderKey     = "Authorization"
	IDOfContextKey       = "id"
	UsernameOfContextKey = "username"
	EmailOfSubject       = "Todo List Backup"
)
