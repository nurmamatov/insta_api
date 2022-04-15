package models

type CreateUserReq struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Phone     string `json:"phonoe"`
	Email     string `json:"email"`
	Gender    bool   `json:"gender"`
	Possword  string `json:"password"`
	Photo     string `json:"photo"`
}
type LoginReq struct {
	Username string `json:"username"`
	Possword string `json:"password"`
}
type SearchUserReq struct {
	FromUsername string `json:"username"`
}
type GetUserReq struct {
	Username string `json:"username"`
}
type UpdateUserReq struct {
	UserId    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Phone     string `json:"phonoe"`
	Email     string `json:"email"`
	Gender    bool   `json:"gender"`
	Photo     string `json:"photo"`
}
type UpdatePass struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
type DeleteUserReq struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}
type Message struct {
	Message string `json:"message"`
}
type GetUserRes struct {
	UserId    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Phone     string `json:"phonoe"`
	Email     string `json:"email"`
	Photo     string `json:"photo"`
	Gender    bool   `json:"gender"`
	CreatedAt string `json:"created_at"`
}
type SearchRes struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
}
type UserList struct {
	Users []SearchRes `json:"users"`
}
type Err struct {
	Error string `json:"error"`
}
