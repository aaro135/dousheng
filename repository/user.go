package repository

var UserLoginInfo map[string]User

type User struct {
	Id int64 `gorm:column:user_id`
}
