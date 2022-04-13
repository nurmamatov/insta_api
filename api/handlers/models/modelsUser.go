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
	FromUsername string `json:"search"`
}
type GetUserReq struct {
	UserId string `json:"user_id"`
}
type UpdateUserReq struct {
	UserId    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Phone     string `json:"phonoe"`
	Email     string `json:"email"`
	Gender    bool   `json:"gender"`
	Possword  string `json:"password"`
	Photo     string `json:"photo"`
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
type UserList struct {
	Users []GetUserRes `json:"users"`
}
type Err struct {
	Error string `json:"error"`
}
type ListUserssRes struct {
	Posts []GetUserRes `json:"posts"`
}